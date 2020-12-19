package cameras

import (
	domain_cameras "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
)

type camera struct {
	original   domain_cameras.Camera
	projection Matrix
	position   Matrix
}

func createCamera(
	original domain_cameras.Camera,
	projection Matrix,
	position Matrix,
) Camera {
	out := camera{
		original:   original,
		projection: projection,
		position:   position,
	}

	return &out
}

// Original returns the original camera
func (obj *camera) Original() domain_cameras.Camera {
	return obj.original
}

// Position returns the position
func (obj *camera) Position() Matrix {
	return obj.position
}

// Projection returns the projection
func (obj *camera) Projection() Matrix {
	return obj.projection
}

// Render renders the camera
func (obj *camera) Render() error {
	return nil
}
