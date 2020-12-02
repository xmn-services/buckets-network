package accesses

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
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

// Adapter returns the accesses adapter
type Adapter interface {
	ToJSON(accesses Accesses) *JSONAccesses
	ToAccesses(ins *JSONAccesses) (Accesses, error)
}

// Factory represents an accesses factory
type Factory interface {
	Create() (Accesses, error)
}

// Builder represents an accesses builder
type Builder interface {
	Create() Builder
	WithList(lst []hash.Hash) Builder
	Now() (Accesses, error)
}

// Accesses represents accesses
type Accesses interface {
	All() []hash.Hash
	Add(bucket buckets.Bucket) error
	Fetch(bucket hash.Hash) (*hash.Hash, error)
	Delete(bucket hash.Hash) error
}
