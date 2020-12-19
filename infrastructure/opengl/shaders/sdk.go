package shaders

import (
	domain_shader "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders/shader"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	shaderBuilder := shader.NewBuilder()
	return createBuilder(shaderBuilder)
}

// Builder represents a shaders builder
type Builder interface {
	Create() Builder
	WithShaders(shaders []domain_shader.Shader) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	All() []shader.Shader
}
