package shader

import (
	domain_shader "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders/shader"
)

type shader struct {
	shader     domain_shader.Shader
	identifier uint32
}

func createShader(
	sh domain_shader.Shader,
	identifier uint32,
) Shader {
	out := shader{
		shader:     sh,
		identifier: identifier,
	}

	return &out
}

// Shader returns the shader
func (obj *shader) Shader() domain_shader.Shader {
	return obj.shader
}

// Identifier returns the identifier
func (obj *shader) Identifier() uint32 {
	return obj.identifier
}
