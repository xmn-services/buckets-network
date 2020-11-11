package bucket

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	pkAdapter := encryption.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, pkAdapter, immutableBuilder)
}

// Adapter represents the bucket adapter
type Adapter interface {
	ToJSON(ins Bucket) *JSONBucket
	ToBucket(js *JSONBucket) (Bucket, error)
}

// Builder represents a bucket builder
type Builder interface {
	Create() Builder
	WithBucket(bucket buckets.Bucket) Builder
	WithAbsolutePath(absolutePath string) Builder
	WithPrivateKey(pk encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bucket, error)
}

// Bucket represents a bucket
type Bucket interface {
	entities.Immutable
	Bucket() buckets.Bucket
	AbsolutePath() string
	PrivateKey() encryption.PrivateKey
}
