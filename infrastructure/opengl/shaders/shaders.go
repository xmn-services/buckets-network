package shaders

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/shaders/shader"
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

// All returns all shaders
func (obj *shaders) All() []shader.Shader {
	return obj.list
}
