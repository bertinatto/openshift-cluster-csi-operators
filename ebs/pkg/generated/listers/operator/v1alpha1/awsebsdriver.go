// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/aws-ebs-csi-driver-operator/pkg/apis/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSEBSDriverLister helps list AWSEBSDrivers.
type AWSEBSDriverLister interface {
	// List lists all AWSEBSDrivers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AWSEBSDriver, err error)
	// Get retrieves the AWSEBSDriver from the index for a given name.
	Get(name string) (*v1alpha1.AWSEBSDriver, error)
	AWSEBSDriverListerExpansion
}

// aWSEBSDriverLister implements the AWSEBSDriverLister interface.
type aWSEBSDriverLister struct {
	indexer cache.Indexer
}

// NewAWSEBSDriverLister returns a new AWSEBSDriverLister.
func NewAWSEBSDriverLister(indexer cache.Indexer) AWSEBSDriverLister {
	return &aWSEBSDriverLister{indexer: indexer}
}

// List lists all AWSEBSDrivers in the indexer.
func (s *aWSEBSDriverLister) List(selector labels.Selector) (ret []*v1alpha1.AWSEBSDriver, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSEBSDriver))
	})
	return ret, err
}

// Get retrieves the AWSEBSDriver from the index for a given name.
func (s *aWSEBSDriverLister) Get(name string) (*v1alpha1.AWSEBSDriver, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsebsdriver"), name)
	}
	return obj.(*v1alpha1.AWSEBSDriver), nil
}
