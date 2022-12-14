// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// This file is part of ewm-cloud-robotics
// (see https://github.com/SAP/ewm-cloud-robotics).
//
// This file is licensed under the Apache Software License, v. 2 except as noted
// otherwise in the LICENSE file (https://github.com/SAP/ewm-cloud-robotics/blob/main/LICENSE)
//

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	missionv1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/apis/mission/v1alpha1"
	versioned "github.com/SAP/ewm-cloud-robotics/go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/SAP/ewm-cloud-robotics/go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/SAP/ewm-cloud-robotics/go/pkg/client/listers/mission/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MissionInformer provides access to a shared informer and lister for
// Missions.
type MissionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MissionLister
}

type missionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMissionInformer constructs a new informer for Mission type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMissionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMissionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMissionInformer constructs a new informer for Mission type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMissionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MissionV1alpha1().Missions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MissionV1alpha1().Missions(namespace).Watch(context.TODO(), options)
			},
		},
		&missionv1alpha1.Mission{},
		resyncPeriod,
		indexers,
	)
}

func (f *missionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMissionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *missionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&missionv1alpha1.Mission{}, f.defaultInformer)
}

func (f *missionInformer) Lister() v1alpha1.MissionLister {
	return v1alpha1.NewMissionLister(f.Informer().GetIndexer())
}
