/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClickHouseVersionsGetter has a method to return a ClickHouseVersionInterface.
// A group's client should implement this interface.
type ClickHouseVersionsGetter interface {
	ClickHouseVersions() ClickHouseVersionInterface
}

// ClickHouseVersionInterface has methods to work with ClickHouseVersion resources.
type ClickHouseVersionInterface interface {
	Create(ctx context.Context, clickHouseVersion *v1alpha1.ClickHouseVersion, opts v1.CreateOptions) (*v1alpha1.ClickHouseVersion, error)
	Update(ctx context.Context, clickHouseVersion *v1alpha1.ClickHouseVersion, opts v1.UpdateOptions) (*v1alpha1.ClickHouseVersion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClickHouseVersion, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClickHouseVersionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClickHouseVersion, err error)
	ClickHouseVersionExpansion
}

// clickHouseVersions implements ClickHouseVersionInterface
type clickHouseVersions struct {
	client rest.Interface
}

// newClickHouseVersions returns a ClickHouseVersions
func newClickHouseVersions(c *CatalogV1alpha1Client) *clickHouseVersions {
	return &clickHouseVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the clickHouseVersion, and returns the corresponding clickHouseVersion object, and an error if there is any.
func (c *clickHouseVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClickHouseVersion, err error) {
	result = &v1alpha1.ClickHouseVersion{}
	err = c.client.Get().
		Resource("clickhouseversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClickHouseVersions that match those selectors.
func (c *clickHouseVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClickHouseVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClickHouseVersionList{}
	err = c.client.Get().
		Resource("clickhouseversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clickHouseVersions.
func (c *clickHouseVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clickhouseversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clickHouseVersion and creates it.  Returns the server's representation of the clickHouseVersion, and an error, if there is any.
func (c *clickHouseVersions) Create(ctx context.Context, clickHouseVersion *v1alpha1.ClickHouseVersion, opts v1.CreateOptions) (result *v1alpha1.ClickHouseVersion, err error) {
	result = &v1alpha1.ClickHouseVersion{}
	err = c.client.Post().
		Resource("clickhouseversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clickHouseVersion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clickHouseVersion and updates it. Returns the server's representation of the clickHouseVersion, and an error, if there is any.
func (c *clickHouseVersions) Update(ctx context.Context, clickHouseVersion *v1alpha1.ClickHouseVersion, opts v1.UpdateOptions) (result *v1alpha1.ClickHouseVersion, err error) {
	result = &v1alpha1.ClickHouseVersion{}
	err = c.client.Put().
		Resource("clickhouseversions").
		Name(clickHouseVersion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clickHouseVersion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clickHouseVersion and deletes it. Returns an error if one occurs.
func (c *clickHouseVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clickhouseversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clickHouseVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clickhouseversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clickHouseVersion.
func (c *clickHouseVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClickHouseVersion, err error) {
	result = &v1alpha1.ClickHouseVersion{}
	err = c.client.Patch(pt).
		Resource("clickhouseversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
