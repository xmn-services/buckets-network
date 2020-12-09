package rectangles

import "github.com/xmn-services/buckets-network/domain/memory/worlds/math"

type rectangle struct {
	pos math.Vec2
	dim math.Vec2
}

func createRectangle(
	pos math.Vec2,
	dim math.Vec2,
) Rectangle {
	out := rectangle{
		pos: pos,
		dim: dim,
	}

	return &out
}

// Position returns the position
func (obj *rectangle) Position() math.Vec2 {
	return obj.pos
}

// Dimension returns the dimension
func (obj *rectangle) Dimension() math.Vec2 {
	return obj.dim
}
