package follows

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Factory represents a follow factory
type Factory interface {
	Create() Follow
}

// Follow represents a buckets follow
type Follow interface {
	entities.Mutable
	Buckets() []buckets.Bucket
	Requests() []hash.Hash
	Add(bucket buckets.Bucket) error
	Request(hash hash.Hash) error
}
