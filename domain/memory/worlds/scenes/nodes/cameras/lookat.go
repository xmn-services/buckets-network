package cameras

import "github.com/xmn-services/buckets-network/domain/memory/worlds/math"

type lookAt struct {
	variable string
	eye      math.Vec3
	center   math.Vec3
	up       math.Vec3
}

func createLookAt(
	variable string,
	eye math.Vec3,
	center math.Vec3,
	up math.Vec3,
) LookAt {
	out := lookAt{
		variable: variable,
		eye:      eye,
		center:   center,
		up:       up,
	}

	return &out
}

// Variable returns the variable
func (obj *lookAt) Variable() string {
	return obj.variable
}

// Eye returns the eye of the camera
func (obj *lookAt) Eye() math.Vec3 {
	return obj.eye
}

// Center returns the center of the camera
func (obj *lookAt) Center() math.Vec3 {
	return obj.center
}

// Up returns the up of the camera
func (obj *lookAt) Up() math.Vec3 {
	return obj.up
}
