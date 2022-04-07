/*
Copyright 2022.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	crdtryv1 "kubebuilder.test/crdtry/api/crdtry/v1"
)

// FakeLpxpods implements LpxpodInterface
type FakeLpxpods struct {
	Fake *FakeCrdtryV1
	ns   string
}

var lpxpodsResource = schema.GroupVersionResource{Group: "crdtry.kubebuilder.test", Version: "v1", Resource: "lpxpods"}

var lpxpodsKind = schema.GroupVersionKind{Group: "crdtry.kubebuilder.test", Version: "v1", Kind: "Lpxpod"}

// Get takes name of the lpxpod, and returns the corresponding lpxpod object, and an error if there is any.
func (c *FakeLpxpods) Get(ctx context.Context, name string, options v1.GetOptions) (result *crdtryv1.Lpxpod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lpxpodsResource, c.ns, name), &crdtryv1.Lpxpod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crdtryv1.Lpxpod), err
}

// List takes label and field selectors, and returns the list of Lpxpods that match those selectors.
func (c *FakeLpxpods) List(ctx context.Context, opts v1.ListOptions) (result *crdtryv1.LpxpodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lpxpodsResource, lpxpodsKind, c.ns, opts), &crdtryv1.LpxpodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &crdtryv1.LpxpodList{ListMeta: obj.(*crdtryv1.LpxpodList).ListMeta}
	for _, item := range obj.(*crdtryv1.LpxpodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lpxpods.
func (c *FakeLpxpods) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lpxpodsResource, c.ns, opts))

}

// Create takes the representation of a lpxpod and creates it.  Returns the server's representation of the lpxpod, and an error, if there is any.
func (c *FakeLpxpods) Create(ctx context.Context, lpxpod *crdtryv1.Lpxpod, opts v1.CreateOptions) (result *crdtryv1.Lpxpod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lpxpodsResource, c.ns, lpxpod), &crdtryv1.Lpxpod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crdtryv1.Lpxpod), err
}

// Update takes the representation of a lpxpod and updates it. Returns the server's representation of the lpxpod, and an error, if there is any.
func (c *FakeLpxpods) Update(ctx context.Context, lpxpod *crdtryv1.Lpxpod, opts v1.UpdateOptions) (result *crdtryv1.Lpxpod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lpxpodsResource, c.ns, lpxpod), &crdtryv1.Lpxpod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crdtryv1.Lpxpod), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLpxpods) UpdateStatus(ctx context.Context, lpxpod *crdtryv1.Lpxpod, opts v1.UpdateOptions) (*crdtryv1.Lpxpod, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lpxpodsResource, "status", c.ns, lpxpod), &crdtryv1.Lpxpod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crdtryv1.Lpxpod), err
}

// Delete takes name of the lpxpod and deletes it. Returns an error if one occurs.
func (c *FakeLpxpods) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(lpxpodsResource, c.ns, name, opts), &crdtryv1.Lpxpod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLpxpods) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lpxpodsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &crdtryv1.LpxpodList{})
	return err
}

// Patch applies the patch and returns the patched lpxpod.
func (c *FakeLpxpods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *crdtryv1.Lpxpod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lpxpodsResource, c.ns, name, pt, data, subresources...), &crdtryv1.Lpxpod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crdtryv1.Lpxpod), err
}
