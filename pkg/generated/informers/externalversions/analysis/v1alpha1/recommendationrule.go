// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	analysisv1alpha1 "github.com/erikwoo/gocrane-api/analysis/v1alpha1"
	versioned "github.com/erikwoo/gocrane-api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/erikwoo/gocrane-api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/erikwoo/gocrane-api/pkg/generated/listers/analysis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RecommendationRuleInformer provides access to a shared informer and lister for
// RecommendationRules.
type RecommendationRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RecommendationRuleLister
}

type recommendationRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRecommendationRuleInformer constructs a new informer for RecommendationRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRecommendationRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRecommendationRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRecommendationRuleInformer constructs a new informer for RecommendationRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRecommendationRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalysisV1alpha1().RecommendationRules().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalysisV1alpha1().RecommendationRules().Watch(context.TODO(), options)
			},
		},
		&analysisv1alpha1.RecommendationRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *recommendationRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRecommendationRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *recommendationRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&analysisv1alpha1.RecommendationRule{}, f.defaultInformer)
}

func (f *recommendationRuleInformer) Lister() v1alpha1.RecommendationRuleLister {
	return v1alpha1.NewRecommendationRuleLister(f.Informer().GetIndexer())
}
