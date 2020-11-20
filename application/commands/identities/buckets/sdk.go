package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files/chunks"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	identity_buckets "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	bucketRepository buckets.Repository,
	identityRepository identities.Repository,
	identityService identities.Service,
	chunkSizeInBytes uint,
	encPKBitrate int,
) Builder {
	hashAdapter := hash.NewAdapter()
	pkFactory := encryption.NewFactory(encPKBitrate)
	chunkBuilder := chunks.NewBuilder()
	fileBuilder := files.NewBuilder()
	bucketBuilder := buckets.NewBuilder()
	identityBucketBuilder := identity_buckets.NewBuilder()

	return createBuilder(
		hashAdapter,
		pkFactory,
		chunkBuilder,
		fileBuilder,
		bucketBuilder,
		bucketRepository,
		identityBucketBuilder,
		identityRepository,
		identityService,
		chunkSizeInBytes,
	)
}

// Builder represents a bucket application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents the bucket application
type Application interface {
	Add(absolutePath string) error
	Delete(absolutePath string) error
	Retrieve(hashStr string) (buckets.Bucket, error)
}
