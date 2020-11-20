package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/file"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	identityService identities.Service,
	fileService file.Service,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
		identityRepository,
		identityService,
		fileService,
	)
}

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
