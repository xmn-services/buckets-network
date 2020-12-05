package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketRepository buckets.Repository,
	contentService contents.Service,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
		identityRepository,
		identityService,
		bucketRepository,
		contentService,
	)
}

// Builder represents a contact application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithBucket(bucketHashStr string) Builder
	Now() (Application, error)
}

// Application represents the bucket application
type Application interface {
	Delete() error
	Retrieve() (buckets.Bucket, error)
	Extract(absolutePath string) error
}
