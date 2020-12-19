package spaces

import "github.com/go-gl/mathgl/mgl32"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a space builder
type Builder interface {
	Create() Builder
	WithInitial(space Space) Builder
	WithPosition(pos mgl32.Vec3) Builder
	WithOrientation(orientation mgl32.Vec4) Builder
	Now() (Space, error)
}

// Space represents the space position and orientation
type Space interface {
	Position() mgl32.Vec3
	Orientation() mgl32.Vec4
	Add(space Space) Space
}
