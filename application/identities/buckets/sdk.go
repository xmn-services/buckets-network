package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/hash"
)

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
	Retrieve(hash hash.Hash) (buckets.Bucket, error)
}
