package vertices

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"

// Builder represents the vertices builder
type Builder interface {
	Create() Builder
	WithVertices(vertices []vertex.Vertex) Builder
	Now() (Vertices, error)
}

// Vertices represents vertices
type Vertices interface {
	All() [][]vertex.Vertex
}
