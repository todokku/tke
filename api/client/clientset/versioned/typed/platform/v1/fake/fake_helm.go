/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeHelms implements HelmInterface
type FakeHelms struct {
	Fake *FakePlatformV1
}

var helmsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "helms"}

var helmsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "Helm"}

// Get takes name of the helm, and returns the corresponding helm object, and an error if there is any.
func (c *FakeHelms) Get(name string, options v1.GetOptions) (result *platformv1.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(helmsResource, name), &platformv1.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Helm), err
}

// List takes label and field selectors, and returns the list of Helms that match those selectors.
func (c *FakeHelms) List(opts v1.ListOptions) (result *platformv1.HelmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(helmsResource, helmsKind, opts), &platformv1.HelmList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.HelmList{ListMeta: obj.(*platformv1.HelmList).ListMeta}
	for _, item := range obj.(*platformv1.HelmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helms.
func (c *FakeHelms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(helmsResource, opts))
}

// Create takes the representation of a helm and creates it.  Returns the server's representation of the helm, and an error, if there is any.
func (c *FakeHelms) Create(helm *platformv1.Helm) (result *platformv1.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(helmsResource, helm), &platformv1.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Helm), err
}

// Update takes the representation of a helm and updates it. Returns the server's representation of the helm, and an error, if there is any.
func (c *FakeHelms) Update(helm *platformv1.Helm) (result *platformv1.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(helmsResource, helm), &platformv1.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Helm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelms) UpdateStatus(helm *platformv1.Helm) (*platformv1.Helm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(helmsResource, "status", helm), &platformv1.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Helm), err
}

// Delete takes name of the helm and deletes it. Returns an error if one occurs.
func (c *FakeHelms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(helmsResource, name), &platformv1.Helm{})
	return err
}

// Patch applies the patch and returns the patched helm.
func (c *FakeHelms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.Helm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(helmsResource, name, pt, data, subresources...), &platformv1.Helm{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.Helm), err
}
