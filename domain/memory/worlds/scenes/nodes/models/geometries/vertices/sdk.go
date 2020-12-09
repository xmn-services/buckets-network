package vertices

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the vertices builder
type Builder interface {
	Create() Builder
	WithVertices(vertices []vertex.Vertex) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Vertices, error)
}

// Vertices represents vertices
type Vertices interface {
	entities.Immutable
	All() [][]vertex.Vertex
}
