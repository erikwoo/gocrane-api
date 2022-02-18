// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EffectiveHorizontalPodAutoscalers returns a EffectiveHorizontalPodAutoscalerInformer.
	EffectiveHorizontalPodAutoscalers() EffectiveHorizontalPodAutoscalerInformer
	// EffectiveVerticalPodAutoscalers returns a EffectiveVerticalPodAutoscalerInformer.
	EffectiveVerticalPodAutoscalers() EffectiveVerticalPodAutoscalerInformer
	// Substitutes returns a SubstituteInformer.
	Substitutes() SubstituteInformer
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

// EffectiveHorizontalPodAutoscalers returns a EffectiveHorizontalPodAutoscalerInformer.
func (v *version) EffectiveHorizontalPodAutoscalers() EffectiveHorizontalPodAutoscalerInformer {
	return &effectiveHorizontalPodAutoscalerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EffectiveVerticalPodAutoscalers returns a EffectiveVerticalPodAutoscalerInformer.
func (v *version) EffectiveVerticalPodAutoscalers() EffectiveVerticalPodAutoscalerInformer {
	return &effectiveVerticalPodAutoscalerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Substitutes returns a SubstituteInformer.
func (v *version) Substitutes() SubstituteInformer {
	return &substituteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
