/*
Copyright 2023 The Kubernetes Authors.

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

	v1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	scheme "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeSnapshotClassesGetter has a method to return a VolumeSnapshotClassInterface.
// A group's client should implement this interface.
type VolumeSnapshotClassesGetter interface {
	VolumeSnapshotClasses() VolumeSnapshotClassInterface
}

// VolumeSnapshotClassInterface has methods to work with VolumeSnapshotClass resources.
type VolumeSnapshotClassInterface interface {
	Create(ctx context.Context, volumeSnapshotClass *v1.VolumeSnapshotClass, opts metav1.CreateOptions) (*v1.VolumeSnapshotClass, error)
	Update(ctx context.Context, volumeSnapshotClass *v1.VolumeSnapshotClass, opts metav1.UpdateOptions) (*v1.VolumeSnapshotClass, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VolumeSnapshotClass, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VolumeSnapshotClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VolumeSnapshotClass, err error)
	VolumeSnapshotClassExpansion
}

// volumeSnapshotClasses implements VolumeSnapshotClassInterface
type volumeSnapshotClasses struct {
	client rest.Interface
}

// newVolumeSnapshotClasses returns a VolumeSnapshotClasses
func newVolumeSnapshotClasses(c *SnapshotV1Client) *volumeSnapshotClasses {
	return &volumeSnapshotClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeSnapshotClass, and returns the corresponding volumeSnapshotClass object, and an error if there is any.
func (c *volumeSnapshotClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VolumeSnapshotClass, err error) {
	result = &v1.VolumeSnapshotClass{}
	err = c.client.Get().
		Resource("volumesnapshotclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeSnapshotClasses that match those selectors.
func (c *volumeSnapshotClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VolumeSnapshotClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VolumeSnapshotClassList{}
	err = c.client.Get().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotClasses.
func (c *volumeSnapshotClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeSnapshotClass and creates it.  Returns the server's representation of the volumeSnapshotClass, and an error, if there is any.
func (c *volumeSnapshotClasses) Create(ctx context.Context, volumeSnapshotClass *v1.VolumeSnapshotClass, opts metav1.CreateOptions) (result *v1.VolumeSnapshotClass, err error) {
	result = &v1.VolumeSnapshotClass{}
	err = c.client.Post().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeSnapshotClass and updates it. Returns the server's representation of the volumeSnapshotClass, and an error, if there is any.
func (c *volumeSnapshotClasses) Update(ctx context.Context, volumeSnapshotClass *v1.VolumeSnapshotClass, opts metav1.UpdateOptions) (result *v1.VolumeSnapshotClass, err error) {
	result = &v1.VolumeSnapshotClass{}
	err = c.client.Put().
		Resource("volumesnapshotclasses").
		Name(volumeSnapshotClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeSnapshotClass and deletes it. Returns an error if one occurs.
func (c *volumeSnapshotClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumesnapshotclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeSnapshotClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumesnapshotclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeSnapshotClass.
func (c *volumeSnapshotClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VolumeSnapshotClass, err error) {
	result = &v1.VolumeSnapshotClass{}
	err = c.client.Patch(pt).
		Resource("volumesnapshotclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
