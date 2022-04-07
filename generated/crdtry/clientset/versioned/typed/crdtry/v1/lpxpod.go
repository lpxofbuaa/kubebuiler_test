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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "kubebuilder.test/crdtry/api/crdtry/v1"
	scheme "kubebuilder.test/crdtry/generated/crdtry/clientset/versioned/scheme"
)

// LpxpodsGetter has a method to return a LpxpodInterface.
// A group's client should implement this interface.
type LpxpodsGetter interface {
	Lpxpods(namespace string) LpxpodInterface
}

// LpxpodInterface has methods to work with Lpxpod resources.
type LpxpodInterface interface {
	Create(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.CreateOptions) (*v1.Lpxpod, error)
	Update(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.UpdateOptions) (*v1.Lpxpod, error)
	UpdateStatus(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.UpdateOptions) (*v1.Lpxpod, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Lpxpod, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LpxpodList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Lpxpod, err error)
	LpxpodExpansion
}

// lpxpods implements LpxpodInterface
type lpxpods struct {
	client rest.Interface
	ns     string
}

// newLpxpods returns a Lpxpods
func newLpxpods(c *CrdtryV1Client, namespace string) *lpxpods {
	return &lpxpods{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lpxpod, and returns the corresponding lpxpod object, and an error if there is any.
func (c *lpxpods) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Lpxpod, err error) {
	result = &v1.Lpxpod{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lpxpods").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Lpxpods that match those selectors.
func (c *lpxpods) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LpxpodList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LpxpodList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lpxpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lpxpods.
func (c *lpxpods) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lpxpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lpxpod and creates it.  Returns the server's representation of the lpxpod, and an error, if there is any.
func (c *lpxpods) Create(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.CreateOptions) (result *v1.Lpxpod, err error) {
	result = &v1.Lpxpod{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lpxpods").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lpxpod).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lpxpod and updates it. Returns the server's representation of the lpxpod, and an error, if there is any.
func (c *lpxpods) Update(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.UpdateOptions) (result *v1.Lpxpod, err error) {
	result = &v1.Lpxpod{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lpxpods").
		Name(lpxpod.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lpxpod).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lpxpods) UpdateStatus(ctx context.Context, lpxpod *v1.Lpxpod, opts metav1.UpdateOptions) (result *v1.Lpxpod, err error) {
	result = &v1.Lpxpod{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lpxpods").
		Name(lpxpod.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lpxpod).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lpxpod and deletes it. Returns an error if one occurs.
func (c *lpxpods) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lpxpods").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lpxpods) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lpxpods").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lpxpod.
func (c *lpxpods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Lpxpod, err error) {
	result = &v1.Lpxpod{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lpxpods").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
