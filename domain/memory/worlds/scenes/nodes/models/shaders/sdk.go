package shaders

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"

// Builder represents the shaders builder
type Builder interface {
	Create() Builder
	WithShaders(shaders []shader.Shader) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	All() []shader.Shader
}
