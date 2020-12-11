package layer

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type layer struct {
	layer   hash.Hash
	shaders shaders.Shaders
}

func createLayer(
	hsh hash.Hash,
	shaders shaders.Shaders,
) Layer {
	out := layer{
		layer:   hsh,
		shaders: shaders,
	}

	return &out
}

// Layer returns the layer hash
func (obj *layer) Layer() hash.Hash {
	return obj.layer
}

// Compiled returns the compiled shaders
func (obj *layer) Compiled() shaders.Shaders {
	return obj.shaders
}
