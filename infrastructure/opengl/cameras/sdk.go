package cameras

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	domain_cameras "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

const glStrVarPattern = "%s\x00"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a camera builder
type Builder interface {
	Create() Builder
	WithProgram(prog programs.Program) Builder
	WithCamera(camera cameras.Camera) Builder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	Original() domain_cameras.Camera
	Position() Matrix
	Projection() Matrix
	Render() error
}

// Matrix represents a camera matrix
type Matrix interface {
	UniformVariable() int32
	Value() mgl32.Mat4
}
