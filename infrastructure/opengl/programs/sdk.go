package programs

import (
	domain_shaders "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	shadersBuilder := shaders.NewBuilder()
	return createBuilder(shadersBuilder)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithShaders(shaders domain_shaders.Shaders) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Shaders() shaders.Shaders
	Identifier() uint32
}
