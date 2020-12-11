package materials

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials/material"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents materials builder
type Builder interface {
	Create() Builder
	WithMaterials(materials []material.Material) Builder
	Now() (Materials, error)
}

// Materials represents materials
type Materials interface {
	All() []material.Material
	CompiledShaders() []shader.Shader
}
