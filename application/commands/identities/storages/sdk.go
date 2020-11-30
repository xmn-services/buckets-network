package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

const bucketDoesNotExistsErr = "the bucket (hash: %s) does not exists"

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	bucketRepository buckets.Repository,
	contentService contents.Service,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
		identityRepository,
		bucketRepository,
		contentService,
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
	Save(bucketHashStr string, chunk []byte) error
	Delete(bucketHashStr string, chunkHashStr string) error
	DeleteAll(bucketHashStr string) error
}
