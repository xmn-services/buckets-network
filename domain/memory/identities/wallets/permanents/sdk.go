package permanents

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/buckets/bucket"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Factory represents a broadcast factory
type Factory interface {
	Create() Permanent
}

// Permanent represents a permanent list of buckets
type Permanent interface {
	entities.Mutable
	All() []bucket.Bucket
	Add(bucket bucket.Bucket) error
	Fetch(absoluteFilePath string) (bucket.Bucket, error)
}
