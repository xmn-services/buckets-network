package material

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type material struct {
	mat    hash.Hash
	layers layers.Layers
}

func createMaterial(
	mat hash.Hash,
	layers layers.Layers,
) Material {
	out := material{
		mat:    mat,
		layers: layers,
	}

	return &out
}

// Material returns the material
func (obj *material) Material() hash.Hash {
	return obj.mat
}

// Compiled returns the compiled layers
func (obj *material) Compiled() layers.Layers {
	return obj.layers
}
