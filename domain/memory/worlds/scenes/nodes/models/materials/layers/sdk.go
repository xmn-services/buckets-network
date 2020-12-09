package layers

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"

// Builder represents the layers builder
type Builder interface {
	Create() Builder
	WithLayers(layers []layer.Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	All() []layer.Layer
}
