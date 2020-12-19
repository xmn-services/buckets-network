package programs

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders"
)

type program struct {
	shaders    shaders.Shaders
	identifier uint32
}

func createProgram(
	shaders shaders.Shaders,
	identifier uint32,
) Program {
	out := program{
		shaders:    shaders,
		identifier: identifier,
	}

	return &out
}

// Shaders return the shaders
func (obj *program) Shaders() shaders.Shaders {
	return obj.shaders
}

// Identifier return the identifier
func (obj *program) Identifier() uint32 {
	return obj.identifier
}
