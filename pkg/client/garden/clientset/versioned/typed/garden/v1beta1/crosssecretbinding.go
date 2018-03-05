// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/garden/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CrossSecretBindingsGetter has a method to return a CrossSecretBindingInterface.
// A group's client should implement this interface.
type CrossSecretBindingsGetter interface {
	CrossSecretBindings(namespace string) CrossSecretBindingInterface
}

// CrossSecretBindingInterface has methods to work with CrossSecretBinding resources.
type CrossSecretBindingInterface interface {
	Create(*v1beta1.CrossSecretBinding) (*v1beta1.CrossSecretBinding, error)
	Update(*v1beta1.CrossSecretBinding) (*v1beta1.CrossSecretBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.CrossSecretBinding, error)
	List(opts v1.ListOptions) (*v1beta1.CrossSecretBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CrossSecretBinding, err error)
	CrossSecretBindingExpansion
}

// crossSecretBindings implements CrossSecretBindingInterface
type crossSecretBindings struct {
	client rest.Interface
	ns     string
}

// newCrossSecretBindings returns a CrossSecretBindings
func newCrossSecretBindings(c *GardenV1beta1Client, namespace string) *crossSecretBindings {
	return &crossSecretBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the crossSecretBinding, and returns the corresponding crossSecretBinding object, and an error if there is any.
func (c *crossSecretBindings) Get(name string, options v1.GetOptions) (result *v1beta1.CrossSecretBinding, err error) {
	result = &v1beta1.CrossSecretBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CrossSecretBindings that match those selectors.
func (c *crossSecretBindings) List(opts v1.ListOptions) (result *v1beta1.CrossSecretBindingList, err error) {
	result = &v1beta1.CrossSecretBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested crossSecretBindings.
func (c *crossSecretBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a crossSecretBinding and creates it.  Returns the server's representation of the crossSecretBinding, and an error, if there is any.
func (c *crossSecretBindings) Create(crossSecretBinding *v1beta1.CrossSecretBinding) (result *v1beta1.CrossSecretBinding, err error) {
	result = &v1beta1.CrossSecretBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		Body(crossSecretBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a crossSecretBinding and updates it. Returns the server's representation of the crossSecretBinding, and an error, if there is any.
func (c *crossSecretBindings) Update(crossSecretBinding *v1beta1.CrossSecretBinding) (result *v1beta1.CrossSecretBinding, err error) {
	result = &v1beta1.CrossSecretBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		Name(crossSecretBinding.Name).
		Body(crossSecretBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the crossSecretBinding and deletes it. Returns an error if one occurs.
func (c *crossSecretBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *crossSecretBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crosssecretbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched crossSecretBinding.
func (c *crossSecretBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CrossSecretBinding, err error) {
	result = &v1beta1.CrossSecretBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("crosssecretbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
