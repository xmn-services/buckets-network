package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/file"
)

// Builder represents a storage application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents the storage application
type Application interface {
	Save(file file.File) error
	Delete(fileHashStr string) error
}
