package shaders

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

type shaders struct {
	list []shader.Shader
}

func createShaders(
	list []shader.Shader,
) Shaders {
	out := shaders{
		list: list,
	}

	return &out
}

// All returns all compiled shaders
func (obj *shaders) All() []shader.Shader {
	return obj.list
}
