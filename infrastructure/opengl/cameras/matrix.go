package cameras

import "github.com/go-gl/mathgl/mgl32"

type matrix struct {
	uniformVariable int32
	value           mgl32.Mat4
}

func createMatrix(
	uniformVariable int32,
	value mgl32.Mat4,
) Matrix {
	out := matrix{
		uniformVariable: uniformVariable,
		value:           value,
	}

	return &out
}

// UniformVariable returns the uniform variable
func (obj *matrix) UniformVariable() int32 {
	return obj.uniformVariable
}

// Value returns the value
func (obj *matrix) Value() mgl32.Mat4 {
	return obj.value
}
