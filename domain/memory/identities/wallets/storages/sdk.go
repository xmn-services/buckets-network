package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	filesFactory := files.NewFactory()
	return createBuilder(filesFactory)
}

// Adapter represents the storage adapter
type Adapter interface {
	ToJSON(ins Storage) *JSONStorage
	ToStorage(js *JSONStorage) (Storage, error)
}

// Factory represents the storages factory
type Factory interface {
	Create() (Storage, error)
}

// Builder represents a storages builder
type Builder interface {
	Create() Builder
	WithToDownload(toDownload files.Files) Builder
	WithStored(stored files.Files) Builder
	Now() (Storage, error)
}

// Storage represents a storage instance
type Storage interface {
	ToDownload() files.Files
	Stored() files.Files
}
