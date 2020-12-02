package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/contents"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	contentService contents.Service,
	bucketRepository buckets.Repository,
	bucketService buckets.Service,
	identityRepository identities.Repository,
	identityService identities.Service,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
		contentService,
		bucketRepository,
		bucketService,
		identityRepository,
		identityService,
	)
}

// Builder represents a contact application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithList(listHashStr string) Builder
	WithContact(contactHashStr string) Builder
	Now() (Application, error)
}

// Application represents the bucket application
type Application interface {
	Add(absolutePath string) error
	Delete(bucketHashStr string) error
	Retrieve(bucketHashStr string) (buckets.Bucket, error)
	RetrieveAll() ([]buckets.Bucket, error)
}
