// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
	scheme "github.com/projectcalico/apiserver/pkg/client/clientset_generated/internalclientset/scheme"
)

// ClusterInformationsGetter has a method to return a ClusterInformationInterface.
// A group's client should implement this interface.
type ClusterInformationsGetter interface {
	ClusterInformations() ClusterInformationInterface
}

// ClusterInformationInterface has methods to work with ClusterInformation resources.
type ClusterInformationInterface interface {
	Create(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.CreateOptions) (*projectcalico.ClusterInformation, error)
	Update(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.UpdateOptions) (*projectcalico.ClusterInformation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalico.ClusterInformation, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalico.ClusterInformationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.ClusterInformation, err error)
	ClusterInformationExpansion
}

// clusterInformations implements ClusterInformationInterface
type clusterInformations struct {
	client rest.Interface
}

// newClusterInformations returns a ClusterInformations
func newClusterInformations(c *ProjectcalicoClient) *clusterInformations {
	return &clusterInformations{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterInformation, and returns the corresponding clusterInformation object, and an error if there is any.
func (c *clusterInformations) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.ClusterInformation, err error) {
	result = &projectcalico.ClusterInformation{}
	err = c.client.Get().
		Resource("clusterinformations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterInformations that match those selectors.
func (c *clusterInformations) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.ClusterInformationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &projectcalico.ClusterInformationList{}
	err = c.client.Get().
		Resource("clusterinformations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterInformations.
func (c *clusterInformations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterinformations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterInformation and creates it.  Returns the server's representation of the clusterInformation, and an error, if there is any.
func (c *clusterInformations) Create(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.CreateOptions) (result *projectcalico.ClusterInformation, err error) {
	result = &projectcalico.ClusterInformation{}
	err = c.client.Post().
		Resource("clusterinformations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInformation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterInformation and updates it. Returns the server's representation of the clusterInformation, and an error, if there is any.
func (c *clusterInformations) Update(ctx context.Context, clusterInformation *projectcalico.ClusterInformation, opts v1.UpdateOptions) (result *projectcalico.ClusterInformation, err error) {
	result = &projectcalico.ClusterInformation{}
	err = c.client.Put().
		Resource("clusterinformations").
		Name(clusterInformation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInformation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterInformation and deletes it. Returns an error if one occurs.
func (c *clusterInformations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterinformations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterInformations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterinformations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterInformation.
func (c *clusterInformations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.ClusterInformation, err error) {
	result = &projectcalico.ClusterInformation{}
	err = c.client.Patch(pt).
		Resource("clusterinformations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
