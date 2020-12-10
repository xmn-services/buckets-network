package nodes

import "github.com/xmn-services/buckets-network/domain/memory/worlds/math"

type space struct {
	pos   math.Vec3
	right math.Vec3
	up    math.Vec3
}

func createSpace(
	pos math.Vec3,
	right math.Vec3,
	up math.Vec3,
) Space {
	out := space{
		pos:   pos,
		right: right,
		up:    up,
	}

	return &out
}

// Position returns the position
func (obj *space) Position() math.Vec3 {
	return obj.pos
}

// Right returns the right
func (obj *space) Right() math.Vec3 {
	return obj.right
}

// Up returns the up
func (obj *space) Up() math.Vec3 {
	return obj.up
}
