// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
)

// FelixConfigurationLister helps list FelixConfigurations.
// All objects returned here must be treated as read-only.
type FelixConfigurationLister interface {
	// List lists all FelixConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectcalico.FelixConfiguration, err error)
	// Get retrieves the FelixConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*projectcalico.FelixConfiguration, error)
	FelixConfigurationListerExpansion
}

// felixConfigurationLister implements the FelixConfigurationLister interface.
type felixConfigurationLister struct {
	indexer cache.Indexer
}

// NewFelixConfigurationLister returns a new FelixConfigurationLister.
func NewFelixConfigurationLister(indexer cache.Indexer) FelixConfigurationLister {
	return &felixConfigurationLister{indexer: indexer}
}

// List lists all FelixConfigurations in the indexer.
func (s *felixConfigurationLister) List(selector labels.Selector) (ret []*projectcalico.FelixConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*projectcalico.FelixConfiguration))
	})
	return ret, err
}

// Get retrieves the FelixConfiguration from the index for a given name.
func (s *felixConfigurationLister) Get(name string) (*projectcalico.FelixConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(projectcalico.Resource("felixconfiguration"), name)
	}
	return obj.(*projectcalico.FelixConfiguration), nil
}
