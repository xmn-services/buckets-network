package buckets

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
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
	hashAdapter := hash.NewAdapter()
	mutableBuilder := entities.NewMutableBuilder()
	return createBuilder(hashAdapter, mutableBuilder)
}

// Adapter represents the buckets adapter
type Adapter interface {
	ToJSON(ins Buckets) *JSONBuckets
	ToBuckets(js *JSONBuckets) (Buckets, error)
}

// Factory represents a buckets factory
type Factory interface {
	Create() (Buckets, error)
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
	Delete(hash hash.Hash) error
}
