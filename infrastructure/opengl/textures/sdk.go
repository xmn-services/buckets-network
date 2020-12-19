package textures

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a texture builder
type Builder interface {
	Create() Builder
	WithTexture(tex textures.Texture) Builder
	Now() (Texture, error)
}

// Texture represents a texture
type Texture interface {
	Texture() textures.Texture
	Identifier() uint32
}
