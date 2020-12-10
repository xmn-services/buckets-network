package renders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	mutableBuilder := entities.NewMutableBuilder()
	return createBuilder(hashAdapter, mutableBuilder)
}

// Factory represents a renders factory
type Factory interface {
	Create() (Renders, error)
}

// Builder represents the renders builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithRenders(renders []render.Render) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Renders, error)
}

// Renders represents renders
type Renders interface {
	entities.Mutable
	All() []render.Render
}
