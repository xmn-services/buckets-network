package vertices

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"
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

// Factory represents a vertices factory
type Factory interface {
	Create() (Vertices, error)
}

// Builder represents the vertices builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithVertices(vertices []vertex.Vertex) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Vertices, error)
}

// Vertices represents vertices
type Vertices interface {
	entities.Mutable
	All() []vertex.Vertex
}
