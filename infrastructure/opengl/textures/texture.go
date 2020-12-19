package textures

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"

type texture struct {
	tex        textures.Texture
	identifier uint32
}

func createTexture(
	tex textures.Texture,
	identifier uint32,
) Texture {
	out := texture{
		tex:        tex,
		identifier: identifier,
	}

	return &out
}

// Texture returns the texture
func (obj *texture) Texture() textures.Texture {
	return obj.tex
}

// Identifier returns the identifier
func (obj *texture) Identifier() uint32 {
	return obj.identifier
}
