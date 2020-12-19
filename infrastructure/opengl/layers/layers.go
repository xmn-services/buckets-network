package layers

import "github.com/xmn-services/buckets-network/infrastructure/opengl/layers/layer"

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

// All returns all layers
func (obj *layers) All() []layer.Layer {
	return obj.list
}
