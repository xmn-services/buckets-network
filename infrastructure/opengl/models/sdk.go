package models

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

const float32SizeInBytes = 32 / 8

const glStrPattern = "%s\x00"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	programBuilder := programs.NewBuilder()
	materialBuilder := materials.NewBuilder()
	return createBuilder(programBuilder, materialBuilder)
}

// Builder represents a model builder
type Builder interface {
	Create() Builder
	WithModel(model models.Model) Builder
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
	Render(camera cameras.Camera, space spaces.Space) error
}

// Type represents the model type
type Type interface {
	IsTriangle() bool
}
