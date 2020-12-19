package shaders

import (
	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders"
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
	WithShaders(shaders domain_shaders.Shaders) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	All() []shader.Shader
}
