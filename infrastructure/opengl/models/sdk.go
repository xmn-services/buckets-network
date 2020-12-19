package models

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

const float32SizeInBytes = 32 / 8

const glStrPattern = "%s\x00"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	materialBuilder := materials.NewBuilder()
	return createBuilder(materialBuilder)
}

// Builder represents a model builder
type Builder interface {
	Create() Builder
	WithModel(model models.Model) Builder
	WithProgram(prog programs.Program) Builder
	Now() (Model, error)
}

// Model represents a model
type Model interface {
	Model() models.Model
	Type() Type
	VAO() uint32
	VertexAmount() int32
	UniformVariable() int32
	Program() programs.Program
	Material() materials.Material
	Render(pos mgl32.Vec3, orientation mgl32.Vec4) error
}

// Type represents the model type
type Type interface {
	IsTriangle() bool
}
