// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/aws-ebs-csi-driver-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AWSEBSDrivers returns a AWSEBSDriverInformer.
	AWSEBSDrivers() AWSEBSDriverInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AWSEBSDrivers returns a AWSEBSDriverInformer.
func (v *version) AWSEBSDrivers() AWSEBSDriverInformer {
	return &aWSEBSDriverInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
