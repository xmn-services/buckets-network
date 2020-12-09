package layers

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the layers builder
type Builder interface {
	Create() Builder
	WithLayers(layers []layer.Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	entities.Immutable
	All() []layer.Layer
}
