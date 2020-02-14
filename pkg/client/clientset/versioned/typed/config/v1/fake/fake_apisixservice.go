/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	configv1 "github.com/gxthrj/apisix-ingress-types/pkg/apis/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApisixServices implements ApisixServiceInterface
type FakeApisixServices struct {
	Fake *FakeApisixV1
	ns   string
}

var apisixservicesResource = schema.GroupVersionResource{Group: "apisix.apache.org", Version: "v1", Resource: "apisixservices"}

var apisixservicesKind = schema.GroupVersionKind{Group: "apisix.apache.org", Version: "v1", Kind: "ApisixService"}

// Get takes name of the apisixService, and returns the corresponding apisixService object, and an error if there is any.
func (c *FakeApisixServices) Get(name string, options v1.GetOptions) (result *configv1.ApisixService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apisixservicesResource, c.ns, name), &configv1.ApisixService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ApisixService), err
}

// List takes label and field selectors, and returns the list of ApisixServices that match those selectors.
func (c *FakeApisixServices) List(opts v1.ListOptions) (result *configv1.ApisixServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apisixservicesResource, apisixservicesKind, c.ns, opts), &configv1.ApisixServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.ApisixServiceList{ListMeta: obj.(*configv1.ApisixServiceList).ListMeta}
	for _, item := range obj.(*configv1.ApisixServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apisixServices.
func (c *FakeApisixServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apisixservicesResource, c.ns, opts))

}

// Create takes the representation of a apisixService and creates it.  Returns the server's representation of the apisixService, and an error, if there is any.
func (c *FakeApisixServices) Create(apisixService *configv1.ApisixService) (result *configv1.ApisixService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apisixservicesResource, c.ns, apisixService), &configv1.ApisixService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ApisixService), err
}

// Update takes the representation of a apisixService and updates it. Returns the server's representation of the apisixService, and an error, if there is any.
func (c *FakeApisixServices) Update(apisixService *configv1.ApisixService) (result *configv1.ApisixService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apisixservicesResource, c.ns, apisixService), &configv1.ApisixService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ApisixService), err
}

// Delete takes name of the apisixService and deletes it. Returns an error if one occurs.
func (c *FakeApisixServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apisixservicesResource, c.ns, name), &configv1.ApisixService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApisixServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apisixservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &configv1.ApisixServiceList{})
	return err
}

// Patch applies the patch and returns the patched apisixService.
func (c *FakeApisixServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configv1.ApisixService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apisixservicesResource, c.ns, name, pt, data, subresources...), &configv1.ApisixService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.ApisixService), err
}
