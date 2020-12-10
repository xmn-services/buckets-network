package renders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type renders struct {
	mutable entities.Mutable
	list    []render.Render
}

func createRenders(
	mutable entities.Mutable,
	list []render.Render,
) Renders {
	out := renders{
		mutable: mutable,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *renders) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of render
func (obj *renders) All() []render.Render {
	return obj.list
}

// CreatedOn returns the creation time
func (obj *renders) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *renders) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}
