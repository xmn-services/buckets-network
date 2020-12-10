package geometries

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	verticesFactory := vertices.NewFactory()
	return createBuilder(hashAdapter, immutableBuilder, verticesFactory)
}

// Builder represents the geometry builder
type Builder interface {
	Create() Builder
	WithShaders(shaders shaders.Shaders) Builder
	WithVertices(vertices vertices.Vertices) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Geometry, error)
}

// Geometry reporesents a geometry
type Geometry interface {
	entities.Immutable
	Shaders() shaders.Shaders
	Vertices() vertices.Vertices
}
