package shader

import (
	domain "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	builder := NewBuilder()
	return createApplication(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Application represents a shader application
type Application interface {
	Compile(shader domain.Shader) (Shader, error)
}

// Builder represents a shader builder
type Builder interface {
	Create() Builder
	WithShader(shader hash.Hash) Builder
	WithIdentifier(identifier uint32) Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	Shader() hash.Hash
	Identifier() uint32
}
