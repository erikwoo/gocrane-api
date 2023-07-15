// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/erikwoo/gocrane-api/topology/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeResourceTopologyLister helps list NodeResourceTopologies.
// All objects returned here must be treated as read-only.
type NodeResourceTopologyLister interface {
	// List lists all NodeResourceTopologies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeResourceTopology, err error)
	// Get retrieves the NodeResourceTopology from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeResourceTopology, error)
	NodeResourceTopologyListerExpansion
}

// nodeResourceTopologyLister implements the NodeResourceTopologyLister interface.
type nodeResourceTopologyLister struct {
	indexer cache.Indexer
}

// NewNodeResourceTopologyLister returns a new NodeResourceTopologyLister.
func NewNodeResourceTopologyLister(indexer cache.Indexer) NodeResourceTopologyLister {
	return &nodeResourceTopologyLister{indexer: indexer}
}

// List lists all NodeResourceTopologies in the indexer.
func (s *nodeResourceTopologyLister) List(selector labels.Selector) (ret []*v1alpha1.NodeResourceTopology, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeResourceTopology))
	})
	return ret, err
}

// Get retrieves the NodeResourceTopology from the index for a given name.
func (s *nodeResourceTopologyLister) Get(name string) (*v1alpha1.NodeResourceTopology, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("noderesourcetopology"), name)
	}
	return obj.(*v1alpha1.NodeResourceTopology), nil
}
