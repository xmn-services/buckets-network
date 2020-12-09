package geometries

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the geometry builder
type Builder interface {
	Create() Builder
	WithShaders(shaders shaders.Shaders) Builder
	WithVertices(vertices vertices.Vertices) Builder
	Now() (Geometry, error)
}

// Geometry reporesents a geometry
type Geometry interface {
	entities.Immutable
	Shaders() shaders.Shaders
	Vertices() vertices.Vertices
}
