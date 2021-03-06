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
package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedServices implements FederatedServiceInterface
type FakeFederatedServices struct {
	Fake *FakeFederationV1alpha1
	ns   string
}

var federatedservicesResource = schema.GroupVersionResource{Group: "federation.k8s.io", Version: "v1alpha1", Resource: "federatedservices"}

var federatedservicesKind = schema.GroupVersionKind{Group: "federation.k8s.io", Version: "v1alpha1", Kind: "FederatedService"}

// Get takes name of the federatedService, and returns the corresponding federatedService object, and an error if there is any.
func (c *FakeFederatedServices) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedservicesResource, c.ns, name), &v1alpha1.FederatedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedService), err
}

// List takes label and field selectors, and returns the list of FederatedServices that match those selectors.
func (c *FakeFederatedServices) List(opts v1.ListOptions) (result *v1alpha1.FederatedServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedservicesResource, federatedservicesKind, c.ns, opts), &v1alpha1.FederatedServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedServiceList{}
	for _, item := range obj.(*v1alpha1.FederatedServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedServices.
func (c *FakeFederatedServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedservicesResource, c.ns, opts))

}

// Create takes the representation of a federatedService and creates it.  Returns the server's representation of the federatedService, and an error, if there is any.
func (c *FakeFederatedServices) Create(federatedService *v1alpha1.FederatedService) (result *v1alpha1.FederatedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedservicesResource, c.ns, federatedService), &v1alpha1.FederatedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedService), err
}

// Update takes the representation of a federatedService and updates it. Returns the server's representation of the federatedService, and an error, if there is any.
func (c *FakeFederatedServices) Update(federatedService *v1alpha1.FederatedService) (result *v1alpha1.FederatedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedservicesResource, c.ns, federatedService), &v1alpha1.FederatedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedServices) UpdateStatus(federatedService *v1alpha1.FederatedService) (*v1alpha1.FederatedService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedservicesResource, "status", c.ns, federatedService), &v1alpha1.FederatedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedService), err
}

// Delete takes name of the federatedService and deletes it. Returns an error if one occurs.
func (c *FakeFederatedServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedservicesResource, c.ns, name), &v1alpha1.FederatedService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedServiceList{})
	return err
}

// Patch applies the patch and returns the patched federatedService.
func (c *FakeFederatedServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedservicesResource, c.ns, name, data, subresources...), &v1alpha1.FederatedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedService), err
}
