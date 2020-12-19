package nodes

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
)

type orientation struct {
	angle     float32
	direction fl32.Vec3
}

func createOrientation(
	angle float32,
	direction fl32.Vec3,
) Orientation {
	out := orientation{
		angle:     angle,
		direction: direction,
	}

	return &out
}

// Angle returns the angle
func (obj *orientation) Angle() float32 {
	return obj.angle
}

// Direction returns the direction
func (obj *orientation) Direction() fl32.Vec3 {
	return obj.direction
}
