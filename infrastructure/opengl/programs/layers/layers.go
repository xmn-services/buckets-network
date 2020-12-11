package layers

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

type layers struct {
	list []layer.Layer
}

func createLayers(
	list []layer.Layer,
) Layers {
	out := layers{
		list: list,
	}

	return &out
}

// All return all layers
func (obj *layers) All() []layer.Layer {
	return obj.list
}

// CompiledShaders returns compiled shaders
func (obj *layers) CompiledShaders() []shader.Shader {
	out := []shader.Shader{}
	for _, oneLayer := range obj.list {
		all := oneLayer.Compiled().All()
		out = append(out, all...)
	}

	return out
}
