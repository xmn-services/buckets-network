package spaces

import "github.com/go-gl/mathgl/mgl32"

type space struct {
	pos         mgl32.Vec3
	orientation mgl32.Vec4
}

func createSpace(
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
) Space {
	out := space{
		pos:         pos,
		orientation: orientation,
	}

	return &out
}

// Position returns the position
func (obj *space) Position() mgl32.Vec3 {
	return obj.pos
}

// Orientation returns the orientation
func (obj *space) Orientation() mgl32.Vec4 {
	return obj.orientation
}

// Add add 2 spaces
func (obj *space) Add(space Space) Space {
	spacePos := space.Position()
	pos := mgl32.Vec3{
		obj.pos[0] + spacePos[0],
		obj.pos[1] + spacePos[1],
		obj.pos[2] + spacePos[2],
	}

	spaceOrientation := space.Orientation()
	orientation := mgl32.Vec4{
		obj.orientation[0] + spaceOrientation[0],
		obj.orientation[1] + spaceOrientation[1],
		obj.orientation[2] + spaceOrientation[2],
		obj.orientation[2] + spaceOrientation[3],
	}

	if orientation[3] > 360 {
		orientation[3] = orientation[3] - 360
	}

	return createSpace(pos, orientation)
}
