package vertex

import "github.com/xmn-services/buckets-network/domain/memory/worlds/math"

type vertex struct {
	pos math.Vec3
	tex math.Vec2
}

func createVertex(
	pos math.Vec3,
	tex math.Vec2,
) Vertex {
	out := vertex{
		pos: pos,
		tex: tex,
	}

	return &out
}

// Position returns the position
func (obj *vertex) Position() math.Vec3 {
	return obj.pos
}

// Texture returns the texture
func (obj *vertex) Texture() math.Vec2 {
	return obj.tex
}
