package buckets

import "github.com/xmn-services/buckets-network/domain/memory/buckets"

// Builder represents a contact application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithContact(contactHashStr string) Builder
	Now() (Application, error)
}

// Application represents the bucket application
type Application interface {
	Add(relativePath string) error
	Delete(relativePath string) error
	Retrieve(hashStr string) (buckets.Bucket, error)
	RetrieveAll() ([]buckets.Bucket, error)
}
