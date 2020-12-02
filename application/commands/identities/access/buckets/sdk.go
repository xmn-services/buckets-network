package buckets

import "github.com/xmn-services/buckets-network/domain/memory/buckets"

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
