// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/erikwoo/gocrane-api/pkg/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/erikwoo/gocrane-api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TimeSeriesPredictionsGetter has a method to return a TimeSeriesPredictionInterface.
// A group's client should implement this interface.
type TimeSeriesPredictionsGetter interface {
	TimeSeriesPredictions(namespace string) TimeSeriesPredictionInterface
}

// TimeSeriesPredictionInterface has methods to work with TimeSeriesPrediction resources.
type TimeSeriesPredictionInterface interface {
	Create(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.CreateOptions) (*v1alpha1.TimeSeriesPrediction, error)
	Update(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (*v1alpha1.TimeSeriesPrediction, error)
	UpdateStatus(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (*v1alpha1.TimeSeriesPrediction, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TimeSeriesPrediction, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TimeSeriesPredictionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TimeSeriesPrediction, err error)
	TimeSeriesPredictionExpansion
}

// timeSeriesPredictions implements TimeSeriesPredictionInterface
type timeSeriesPredictions struct {
	client rest.Interface
	ns     string
}

// newTimeSeriesPredictions returns a TimeSeriesPredictions
func newTimeSeriesPredictions(c *PredictionV1alpha1Client, namespace string) *timeSeriesPredictions {
	return &timeSeriesPredictions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the timeSeriesPrediction, and returns the corresponding timeSeriesPrediction object, and an error if there is any.
func (c *timeSeriesPredictions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	result = &v1alpha1.TimeSeriesPrediction{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TimeSeriesPredictions that match those selectors.
func (c *timeSeriesPredictions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TimeSeriesPredictionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TimeSeriesPredictionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested timeSeriesPredictions.
func (c *timeSeriesPredictions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a timeSeriesPrediction and creates it.  Returns the server's representation of the timeSeriesPrediction, and an error, if there is any.
func (c *timeSeriesPredictions) Create(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.CreateOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	result = &v1alpha1.TimeSeriesPrediction{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(timeSeriesPrediction).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a timeSeriesPrediction and updates it. Returns the server's representation of the timeSeriesPrediction, and an error, if there is any.
func (c *timeSeriesPredictions) Update(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	result = &v1alpha1.TimeSeriesPrediction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		Name(timeSeriesPrediction.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(timeSeriesPrediction).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *timeSeriesPredictions) UpdateStatus(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	result = &v1alpha1.TimeSeriesPrediction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		Name(timeSeriesPrediction.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(timeSeriesPrediction).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the timeSeriesPrediction and deletes it. Returns an error if one occurs.
func (c *timeSeriesPredictions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *timeSeriesPredictions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("timeseriespredictions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched timeSeriesPrediction.
func (c *timeSeriesPredictions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TimeSeriesPrediction, err error) {
	result = &v1alpha1.TimeSeriesPrediction{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("timeseriespredictions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
