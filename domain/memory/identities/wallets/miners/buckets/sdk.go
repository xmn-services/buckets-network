package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
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
	return createBuilder()
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
	WithBuckets(buckets []bucket.Bucket) Builder
	Now() (Buckets, error)
}

// Buckets represents buckets
type Buckets interface {
	All() []bucket.Bucket
	Add(bucket bucket.Bucket) error
	Delete(hash hash.Hash) error
}
