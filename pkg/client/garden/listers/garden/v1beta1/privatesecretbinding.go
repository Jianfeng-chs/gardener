// Code generated by lister-gen. DO NOT EDIT.

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrivateSecretBindingLister helps list PrivateSecretBindings.
type PrivateSecretBindingLister interface {
	// List lists all PrivateSecretBindings in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.PrivateSecretBinding, err error)
	// PrivateSecretBindings returns an object that can list and get PrivateSecretBindings.
	PrivateSecretBindings(namespace string) PrivateSecretBindingNamespaceLister
	PrivateSecretBindingListerExpansion
}

// privateSecretBindingLister implements the PrivateSecretBindingLister interface.
type privateSecretBindingLister struct {
	indexer cache.Indexer
}

// NewPrivateSecretBindingLister returns a new PrivateSecretBindingLister.
func NewPrivateSecretBindingLister(indexer cache.Indexer) PrivateSecretBindingLister {
	return &privateSecretBindingLister{indexer: indexer}
}

// List lists all PrivateSecretBindings in the indexer.
func (s *privateSecretBindingLister) List(selector labels.Selector) (ret []*v1beta1.PrivateSecretBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PrivateSecretBinding))
	})
	return ret, err
}

// PrivateSecretBindings returns an object that can list and get PrivateSecretBindings.
func (s *privateSecretBindingLister) PrivateSecretBindings(namespace string) PrivateSecretBindingNamespaceLister {
	return privateSecretBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateSecretBindingNamespaceLister helps list and get PrivateSecretBindings.
type PrivateSecretBindingNamespaceLister interface {
	// List lists all PrivateSecretBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.PrivateSecretBinding, err error)
	// Get retrieves the PrivateSecretBinding from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.PrivateSecretBinding, error)
	PrivateSecretBindingNamespaceListerExpansion
}

// privateSecretBindingNamespaceLister implements the PrivateSecretBindingNamespaceLister
// interface.
type privateSecretBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateSecretBindings in the indexer for a given namespace.
func (s privateSecretBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.PrivateSecretBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PrivateSecretBinding))
	})
	return ret, err
}

// Get retrieves the PrivateSecretBinding from the indexer for a given namespace and name.
func (s privateSecretBindingNamespaceLister) Get(name string) (*v1beta1.PrivateSecretBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("privatesecretbinding"), name)
	}
	return obj.(*v1beta1.PrivateSecretBinding), nil
}
