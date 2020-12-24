package opengl

import "github.com/go-gl/mathgl/mgl32"

type hudPosition struct {
	vec      mgl32.Vec2
	variable string
}

func createHudPosition(
	vec mgl32.Vec2,
	variable string,
) HudPosition {
	out := hudPosition{
		vec:      vec,
		variable: variable,
	}

	return &out
}

// Vector returns the position vector
func (obj *hudPosition) Vector() mgl32.Vec2 {
	return obj.vec
}

// Variable returns the position variable
func (obj *hudPosition) Variable() string {
	return obj.variable
}
