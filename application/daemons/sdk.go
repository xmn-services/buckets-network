package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application/syncs"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	syncBuilder syncs.Builder,
) Builder {
	return createBuilder(syncBuilder)
}

// Builder represents a daemon application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithWaitPeriod(waitPeriod time.Duration) Builder
	WithAdditionalBucketsPerBlock(additionalBuckets uint) Builder
	Now() (Application, error)
}

// Application represents a daemon application
type Application interface {
	Start() error
	Stop() error
}
