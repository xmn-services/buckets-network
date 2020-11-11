package follows

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// Application represents a follow application
type Application interface {
	Retrieve(bucket hash.Hash) (buckets.Bucket, error)
}
