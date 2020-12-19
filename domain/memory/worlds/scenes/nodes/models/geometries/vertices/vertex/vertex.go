package vertex

import (
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
)

type vertex struct {
	pos fl32.Vec3
	tex fl32.Vec2
}

func createVertex(
	pos fl32.Vec3,
	tex fl32.Vec2,
) Vertex {
	out := vertex{
		pos: pos,
		tex: tex,
	}

	return &out
}

// Position returns the position
func (obj *vertex) Position() fl32.Vec3 {
	return obj.pos
}

// Texture returns the texture
func (obj *vertex) Texture() fl32.Vec2 {
	return obj.tex
}

// String returns the string representation of the vertex
func (obj *vertex) String() string {
	return fmt.Sprintf("pos: %s, tex: %s", obj.Position().String(), obj.Texture().String())
}
