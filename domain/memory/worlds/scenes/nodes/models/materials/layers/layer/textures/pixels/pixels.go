package pixels

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type pixels struct {
	mutable entities.Mutable
	list    []pixel.Pixel
}

func createPixels(
	mutable entities.Mutable,
	list []pixel.Pixel,
) Pixels {
	out := pixels{
		mutable: mutable,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *pixels) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of pixel
func (obj *pixels) All() []pixel.Pixel {
	return obj.list
}

// Amount returns the amount
func (obj *pixels) Amount() uint {
	return uint(len(obj.list))
}

// CreatedOn returns the creation time
func (obj *pixels) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *pixels) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}
