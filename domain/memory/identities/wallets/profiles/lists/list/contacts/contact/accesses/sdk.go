package accesses

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents an accesses factory
type Factory interface {
	Create() Accesses
}

// Accesses represents accesses
type Accesses interface {
	All() []hash.Hash
	Add(bucket buckets.Bucket) error
	Fetch(bucket hash.Hash) (*hash.Hash, error)
	Delete(bucket hash.Hash) error
}
