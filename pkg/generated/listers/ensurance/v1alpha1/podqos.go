// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/erikwoo/gocrane-api/ensurance/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodQOSLister helps list PodQOSs.
// All objects returned here must be treated as read-only.
type PodQOSLister interface {
	// List lists all PodQOSs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PodQOS, err error)
	// Get retrieves the PodQOS from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PodQOS, error)
	PodQOSListerExpansion
}

// podQOSLister implements the PodQOSLister interface.
type podQOSLister struct {
	indexer cache.Indexer
}

// NewPodQOSLister returns a new PodQOSLister.
func NewPodQOSLister(indexer cache.Indexer) PodQOSLister {
	return &podQOSLister{indexer: indexer}
}

// List lists all PodQOSs in the indexer.
func (s *podQOSLister) List(selector labels.Selector) (ret []*v1alpha1.PodQOS, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodQOS))
	})
	return ret, err
}

// Get retrieves the PodQOS from the index for a given name.
func (s *podQOSLister) Get(name string) (*v1alpha1.PodQOS, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podqos"), name)
	}
	return obj.(*v1alpha1.PodQOS), nil
}
