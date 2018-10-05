// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-prune-operator/pkg/apis/prune/v1alpha1"
	scheme "github.com/openshift/cluster-prune-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PruneServicesGetter has a method to return a PruneServiceInterface.
// A group's client should implement this interface.
type PruneServicesGetter interface {
	PruneServices(namespace string) PruneServiceInterface
}

// PruneServiceInterface has methods to work with PruneService resources.
type PruneServiceInterface interface {
	Create(*v1alpha1.PruneService) (*v1alpha1.PruneService, error)
	Update(*v1alpha1.PruneService) (*v1alpha1.PruneService, error)
	UpdateStatus(*v1alpha1.PruneService) (*v1alpha1.PruneService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PruneService, error)
	List(opts v1.ListOptions) (*v1alpha1.PruneServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PruneService, err error)
	PruneServiceExpansion
}

// pruneServices implements PruneServiceInterface
type pruneServices struct {
	client rest.Interface
	ns     string
}

// newPruneServices returns a PruneServices
func newPruneServices(c *PruneV1alpha1Client, namespace string) *pruneServices {
	return &pruneServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pruneService, and returns the corresponding pruneService object, and an error if there is any.
func (c *pruneServices) Get(name string, options v1.GetOptions) (result *v1alpha1.PruneService, err error) {
	result = &v1alpha1.PruneService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pruneservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PruneServices that match those selectors.
func (c *pruneServices) List(opts v1.ListOptions) (result *v1alpha1.PruneServiceList, err error) {
	result = &v1alpha1.PruneServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pruneservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pruneServices.
func (c *pruneServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pruneservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pruneService and creates it.  Returns the server's representation of the pruneService, and an error, if there is any.
func (c *pruneServices) Create(pruneService *v1alpha1.PruneService) (result *v1alpha1.PruneService, err error) {
	result = &v1alpha1.PruneService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pruneservices").
		Body(pruneService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pruneService and updates it. Returns the server's representation of the pruneService, and an error, if there is any.
func (c *pruneServices) Update(pruneService *v1alpha1.PruneService) (result *v1alpha1.PruneService, err error) {
	result = &v1alpha1.PruneService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pruneservices").
		Name(pruneService.Name).
		Body(pruneService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pruneServices) UpdateStatus(pruneService *v1alpha1.PruneService) (result *v1alpha1.PruneService, err error) {
	result = &v1alpha1.PruneService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pruneservices").
		Name(pruneService.Name).
		SubResource("status").
		Body(pruneService).
		Do().
		Into(result)
	return
}

// Delete takes name of the pruneService and deletes it. Returns an error if one occurs.
func (c *pruneServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pruneservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pruneServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pruneservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pruneService.
func (c *pruneServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PruneService, err error) {
	result = &v1alpha1.PruneService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pruneservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}