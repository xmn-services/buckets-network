package rectangles

import "github.com/xmn-services/buckets-network/domain/memory/worlds/math"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a rectangle builder
type Builder interface {
	Create() Builder
	WithPosition(pos math.Vec2) Builder
	WithDimension(dim math.Vec2) Builder
	Now() (Rectangle, error)
}

// Rectangle represents a 2D rectangle
type Rectangle interface {
	Position() math.Vec2
	Dimension() math.Vec2
}
