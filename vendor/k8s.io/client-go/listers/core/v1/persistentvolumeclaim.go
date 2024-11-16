/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// PersistentVolumeClaimLister helps list PersistentVolumeClaims.
// All objects returned here must be treated as read-only.
type PersistentVolumeClaimLister interface {
	// List lists all PersistentVolumeClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	// PersistentVolumeClaims returns an object that can list and get PersistentVolumeClaims.
	PersistentVolumeClaims(namespace string) PersistentVolumeClaimNamespaceLister
	PersistentVolumeClaimListerExpansion
}

// persistentVolumeClaimLister implements the PersistentVolumeClaimLister interface.
type persistentVolumeClaimLister struct {
	listers.ResourceIndexer[*v1.PersistentVolumeClaim]
}

// NewPersistentVolumeClaimLister returns a new PersistentVolumeClaimLister.
func NewPersistentVolumeClaimLister(indexer cache.Indexer) PersistentVolumeClaimLister {
	return &persistentVolumeClaimLister{listers.New[*v1.PersistentVolumeClaim](indexer, v1.Resource("persistentvolumeclaim"))}
}

// PersistentVolumeClaims returns an object that can list and get PersistentVolumeClaims.
func (s *persistentVolumeClaimLister) PersistentVolumeClaims(namespace string) PersistentVolumeClaimNamespaceLister {
	return persistentVolumeClaimNamespaceLister{listers.NewNamespaced[*v1.PersistentVolumeClaim](s.ResourceIndexer, namespace)}
}

// PersistentVolumeClaimNamespaceLister helps list and get PersistentVolumeClaims.
// All objects returned here must be treated as read-only.
type PersistentVolumeClaimNamespaceLister interface {
	// List lists all PersistentVolumeClaims in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	// Get retrieves the PersistentVolumeClaim from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PersistentVolumeClaim, error)
	PersistentVolumeClaimNamespaceListerExpansion
}

// persistentVolumeClaimNamespaceLister implements the PersistentVolumeClaimNamespaceLister
// interface.
type persistentVolumeClaimNamespaceLister struct {
	listers.ResourceIndexer[*v1.PersistentVolumeClaim]
}
