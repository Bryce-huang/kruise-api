/*
Copyright 2022 The Kruise Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	appsv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	versioned "github.com/openkruise/kruise-api/client/clientset/versioned"
	internalinterfaces "github.com/openkruise/kruise-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openkruise/kruise-api/client/listers/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodePodProbeInformer provides access to a shared informer and lister for
// NodePodProbes.
type NodePodProbeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodePodProbeLister
}

type nodePodProbeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodePodProbeInformer constructs a new informer for NodePodProbe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodePodProbeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodePodProbeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodePodProbeInformer constructs a new informer for NodePodProbe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodePodProbeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().NodePodProbes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().NodePodProbes().Watch(context.TODO(), options)
			},
		},
		&appsv1alpha1.NodePodProbe{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodePodProbeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodePodProbeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodePodProbeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha1.NodePodProbe{}, f.defaultInformer)
}

func (f *nodePodProbeInformer) Lister() v1alpha1.NodePodProbeLister {
	return v1alpha1.NewNodePodProbeLister(f.Informer().GetIndexer())
}
