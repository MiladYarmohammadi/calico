// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	projectcalico "github.com/projectcalico/apiserver/pkg/apis/projectcalico"
	internalclientset "github.com/projectcalico/apiserver/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/projectcalico/apiserver/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/projectcalico/apiserver/pkg/client/listers_generated/projectcalico/internalversion"
)

// HostEndpointInformer provides access to a shared informer and lister for
// HostEndpoints.
type HostEndpointInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.HostEndpointLister
}

type hostEndpointInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHostEndpointInformer constructs a new informer for HostEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHostEndpointInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHostEndpointInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHostEndpointInformer constructs a new informer for HostEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHostEndpointInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Projectcalico().HostEndpoints().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Projectcalico().HostEndpoints().Watch(context.TODO(), options)
			},
		},
		&projectcalico.HostEndpoint{},
		resyncPeriod,
		indexers,
	)
}

func (f *hostEndpointInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHostEndpointInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hostEndpointInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalico.HostEndpoint{}, f.defaultInformer)
}

func (f *hostEndpointInformer) Lister() internalversion.HostEndpointLister {
	return internalversion.NewHostEndpointLister(f.Informer().GetIndexer())
}
