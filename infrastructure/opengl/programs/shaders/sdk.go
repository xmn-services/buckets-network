package shaders

import (
	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

// NewApplication creates a new application insance
func NewApplication() Application {
	shaderApplication := shader.NewApplication()
	builder := NewBuilder()
	return createApplication(builder, shaderApplication)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Application represents a shaders application
type Application interface {
	Compile(shaders domain_shaders.Shaders) (Shaders, error)
}

// Builder represents a shaders builder
type Builder interface {
	Create() Builder
	WithCompiledShaders(shaders []shader.Shader) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	All() []shader.Shader
}
