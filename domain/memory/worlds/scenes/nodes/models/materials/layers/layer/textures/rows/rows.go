package rows

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type rows struct {
	mutable entities.Mutable
	list    []pixels.Pixels
}

func createRows(
	mutable entities.Mutable,
	list []pixels.Pixels,
) Rows {
	out := rows{
		mutable: mutable,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *rows) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of pixels
func (obj *rows) All() []pixels.Pixels {
	return obj.list
}

// Dimension returns the dimension
func (obj *rows) Dimension() (uint, uint) {
	width := obj.list[0].Amount()
	height := uint(len(obj.list))
	return width, height
}

// CreatedOn returns the creation time
func (obj *rows) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *rows) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}
