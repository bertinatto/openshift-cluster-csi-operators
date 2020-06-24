package operator

import (
	"context"
	"fmt"
	"time"

	"k8s.io/klog"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
	csicontrollerset "github.com/openshift/library-go/pkg/operator/csi/controllerset"
	goc "github.com/openshift/library-go/pkg/operator/genericoperatorclient"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	clientbuilder "github.com/bertinatto/csi-driver-controller/pkg/builder"

	"github.com/openshift/aws-ebs-csi-driver-operator/pkg/apis/operator/v1alpha1"
	"github.com/openshift/aws-ebs-csi-driver-operator/pkg/generated"
)

const (
	operandName       = "aws-ebs-csi-driver"
	operandNamespace  = "openshift-aws-ebs-csi-driver"
	operatorNamespace = "openshift-aws-ebs-csi-driver-operator"

	resync = 20 * time.Minute
)

func RunOperator(ctx context.Context, controllerConfig *controllercmd.ControllerContext) error {
	// Create clientsets and informers
	cb, err := clientbuilder.NewClientBuilder("")
	if err != nil {
		klog.Fatalf("error creating clients: %v", err)
	}
	ctrlCtx := clientbuilder.NewControllerContext(cb, ctx.Done(), operandNamespace)
	dynamicClient := ctrlCtx.ClientBuilder.DynamicClientOrDie(operandName)
	kubeClient := ctrlCtx.ClientBuilder.KubeClientOrDie(operandName)

	// Create GenericOperatorclient. This is used by controllers created down below
	gvr := v1alpha1.SchemeGroupVersion.WithResource("awsebsdrivers")
	operatorClient, dynamicInformers, err := goc.NewClusterScopedOperatorClient(controllerConfig.KubeConfig, gvr)
	if err != nil {
		return err
	}

	kubeInformersForNamespace := v1helpers.NewKubeInformersForNamespaces(
		kubeClient,
		"",
		operandNamespace,
		operatorNamespace)

	csiControllerSet := csicontrollerset.New(
		operatorClient,
		controllerConfig.EventRecorder,
	).WithLogLevelController().WithManagementStateController(
		operandName,
		false,
	).WithStaticResourcesController(
		"AWSEBSDriverStaticResources",
		kubeClient,
		kubeInformersForNamespace,
		generated.Asset,
		[]string{
			"namespace.yaml",
			"storageclass.yaml",
			"csidriver.yaml",
			"controller_sa.yaml",
			"node_sa.yaml",
			"rbac/provisioner_binding.yaml",
			"rbac/provisioner_role.yaml",
			"rbac/attacher_binding.yaml",
			"rbac/attacher_role.yaml",
			"rbac/privileged_role.yaml",
			"rbac/controller_privileged_binding.yaml",
			"rbac/node_privileged_binding.yaml",
		},
	).WithCSIDriverController(
		"AWSEBSDriverController",
		operandName,
		operandNamespace,
		generated.MustAsset,
		kubeClient,
		ctrlCtx.KubeNamespacedInformerFactory,
		csicontrollerset.WithControllerService("controller.yaml"),
		csicontrollerset.WithNodeService("node.yaml"),
		csicontrollerset.WithCloudCredentials(dynamicClient, "credentials.yaml"),
	)

	if err != nil {
		return err
	}

	klog.Info("Starting the informers")
	go ctrlCtx.KubeNamespacedInformerFactory.Start(ctx.Done())
	go kubeInformersForNamespace.Start(ctx.Done())
	go dynamicInformers.Start(ctx.Done())

	klog.Info("Starting controllerset")
	go csiControllerSet.Run(ctx, 1)

	<-ctx.Done()

	return fmt.Errorf("stopped")
}
