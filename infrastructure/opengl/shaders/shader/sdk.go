package shader

import (
	domain_shader "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a shader builder
type Builder interface {
	Create() Builder
	WithShader(shader domain_shader.Shader) Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	Shader() domain_shader.Shader
	Identifier() uint32
}
