package permanents

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a buckets factory
type Factory interface {
	Create() Buckets
}

// Builder represents a buckets builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithBuckets(buckets []bucket.Bucket) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Buckets, error)
}

// Buckets represents buckets
type Buckets interface {
	entities.Mutable
	All() []bucket.Bucket
	Add(bucket bucket.Bucket) error
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
}
