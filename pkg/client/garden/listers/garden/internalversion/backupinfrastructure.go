// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	garden "github.com/gardener/gardener/pkg/apis/garden"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupInfrastructureLister helps list BackupInfrastructures.
type BackupInfrastructureLister interface {
	// List lists all BackupInfrastructures in the indexer.
	List(selector labels.Selector) (ret []*garden.BackupInfrastructure, err error)
	// BackupInfrastructures returns an object that can list and get BackupInfrastructures.
	BackupInfrastructures(namespace string) BackupInfrastructureNamespaceLister
	BackupInfrastructureListerExpansion
}

// backupInfrastructureLister implements the BackupInfrastructureLister interface.
type backupInfrastructureLister struct {
	indexer cache.Indexer
}

// NewBackupInfrastructureLister returns a new BackupInfrastructureLister.
func NewBackupInfrastructureLister(indexer cache.Indexer) BackupInfrastructureLister {
	return &backupInfrastructureLister{indexer: indexer}
}

// List lists all BackupInfrastructures in the indexer.
func (s *backupInfrastructureLister) List(selector labels.Selector) (ret []*garden.BackupInfrastructure, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*garden.BackupInfrastructure))
	})
	return ret, err
}

// BackupInfrastructures returns an object that can list and get BackupInfrastructures.
func (s *backupInfrastructureLister) BackupInfrastructures(namespace string) BackupInfrastructureNamespaceLister {
	return backupInfrastructureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupInfrastructureNamespaceLister helps list and get BackupInfrastructures.
type BackupInfrastructureNamespaceLister interface {
	// List lists all BackupInfrastructures in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*garden.BackupInfrastructure, err error)
	// Get retrieves the BackupInfrastructure from the indexer for a given namespace and name.
	Get(name string) (*garden.BackupInfrastructure, error)
	BackupInfrastructureNamespaceListerExpansion
}

// backupInfrastructureNamespaceLister implements the BackupInfrastructureNamespaceLister
// interface.
type backupInfrastructureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupInfrastructures in the indexer for a given namespace.
func (s backupInfrastructureNamespaceLister) List(selector labels.Selector) (ret []*garden.BackupInfrastructure, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*garden.BackupInfrastructure))
	})
	return ret, err
}

// Get retrieves the BackupInfrastructure from the indexer for a given namespace and name.
func (s backupInfrastructureNamespaceLister) Get(name string) (*garden.BackupInfrastructure, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(garden.Resource("backupinfrastructure"), name)
	}
	return obj.(*garden.BackupInfrastructure), nil
}