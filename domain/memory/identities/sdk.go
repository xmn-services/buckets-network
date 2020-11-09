package identities

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/identities/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithSeed(seed string) Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithBuckets(buckets []bucket.Bucket) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	entities.Mutable
	Seed() string
	Name() string
	Root() string
	Buckets() buckets.Buckets
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(name string, password string, seed string) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Insert(identity Identity, password string) error
	Update(originalHash hash.Hash, updated Identity, password string, newPassword string) error
	Delete(identity Identity, password string) error
}
