package layers

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type layers struct {
	mutable entities.Mutable
	list    []layer.Layer
}

func createLayers(
	mutable entities.Mutable,
	list []layer.Layer,
) Layers {
	out := layers{
		mutable: mutable,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *layers) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of layers
func (obj *layers) All() []layer.Layer {
	return obj.list
}

// CreatedOn returns the creation time
func (obj *layers) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *layers) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}
