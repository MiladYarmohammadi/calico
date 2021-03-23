// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
)

// HostEndpointLister helps list HostEndpoints.
// All objects returned here must be treated as read-only.
type HostEndpointLister interface {
	// List lists all HostEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.HostEndpoint, err error)
	// Get retrieves the HostEndpoint from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.HostEndpoint, error)
	HostEndpointListerExpansion
}

// hostEndpointLister implements the HostEndpointLister interface.
type hostEndpointLister struct {
	indexer cache.Indexer
}

// NewHostEndpointLister returns a new HostEndpointLister.
func NewHostEndpointLister(indexer cache.Indexer) HostEndpointLister {
	return &hostEndpointLister{indexer: indexer}
}

// List lists all HostEndpoints in the indexer.
func (s *hostEndpointLister) List(selector labels.Selector) (ret []*v3.HostEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.HostEndpoint))
	})
	return ret, err
}

// Get retrieves the HostEndpoint from the index for a given name.
func (s *hostEndpointLister) Get(name string) (*v3.HostEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("hostendpoint"), name)
	}
	return obj.(*v3.HostEndpoint), nil
}
