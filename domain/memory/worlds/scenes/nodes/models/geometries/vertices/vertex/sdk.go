package vertex

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math"
)

// Builder represents a vertex builder
type Builder interface {
	Create() Builder
	WithPosition(pos math.Vec3) Builder
	WithTexture(tex math.Vec2) Builder
	Now() (Vertex, error)
}

// Vertex represents a vertex
type Vertex interface {
	Position() math.Vec3
	Texture() math.Vec2
}
