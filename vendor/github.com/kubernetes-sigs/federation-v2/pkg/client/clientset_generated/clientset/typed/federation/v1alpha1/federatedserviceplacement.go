/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedServicePlacementsGetter has a method to return a FederatedServicePlacementInterface.
// A group's client should implement this interface.
type FederatedServicePlacementsGetter interface {
	FederatedServicePlacements(namespace string) FederatedServicePlacementInterface
}

// FederatedServicePlacementInterface has methods to work with FederatedServicePlacement resources.
type FederatedServicePlacementInterface interface {
	Create(*v1alpha1.FederatedServicePlacement) (*v1alpha1.FederatedServicePlacement, error)
	Update(*v1alpha1.FederatedServicePlacement) (*v1alpha1.FederatedServicePlacement, error)
	UpdateStatus(*v1alpha1.FederatedServicePlacement) (*v1alpha1.FederatedServicePlacement, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedServicePlacement, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedServicePlacementList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServicePlacement, err error)
	FederatedServicePlacementExpansion
}

// federatedServicePlacements implements FederatedServicePlacementInterface
type federatedServicePlacements struct {
	client rest.Interface
	ns     string
}

// newFederatedServicePlacements returns a FederatedServicePlacements
func newFederatedServicePlacements(c *FederationV1alpha1Client, namespace string) *federatedServicePlacements {
	return &federatedServicePlacements{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedServicePlacement, and returns the corresponding federatedServicePlacement object, and an error if there is any.
func (c *federatedServicePlacements) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedServicePlacement, err error) {
	result = &v1alpha1.FederatedServicePlacement{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedServicePlacements that match those selectors.
func (c *federatedServicePlacements) List(opts v1.ListOptions) (result *v1alpha1.FederatedServicePlacementList, err error) {
	result = &v1alpha1.FederatedServicePlacementList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedServicePlacements.
func (c *federatedServicePlacements) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedServicePlacement and creates it.  Returns the server's representation of the federatedServicePlacement, and an error, if there is any.
func (c *federatedServicePlacements) Create(federatedServicePlacement *v1alpha1.FederatedServicePlacement) (result *v1alpha1.FederatedServicePlacement, err error) {
	result = &v1alpha1.FederatedServicePlacement{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		Body(federatedServicePlacement).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedServicePlacement and updates it. Returns the server's representation of the federatedServicePlacement, and an error, if there is any.
func (c *federatedServicePlacements) Update(federatedServicePlacement *v1alpha1.FederatedServicePlacement) (result *v1alpha1.FederatedServicePlacement, err error) {
	result = &v1alpha1.FederatedServicePlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		Name(federatedServicePlacement.Name).
		Body(federatedServicePlacement).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedServicePlacements) UpdateStatus(federatedServicePlacement *v1alpha1.FederatedServicePlacement) (result *v1alpha1.FederatedServicePlacement, err error) {
	result = &v1alpha1.FederatedServicePlacement{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		Name(federatedServicePlacement.Name).
		SubResource("status").
		Body(federatedServicePlacement).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedServicePlacement and deletes it. Returns an error if one occurs.
func (c *federatedServicePlacements) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedServicePlacements) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedServicePlacement.
func (c *federatedServicePlacements) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedServicePlacement, err error) {
	result = &v1alpha1.FederatedServicePlacement{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedserviceplacements").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
