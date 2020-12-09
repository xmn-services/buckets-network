package cameras

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type camera struct {
	immutable entities.Immutable
	viewport  rectangles.Rectangle
	fov       float64
	index     uint
}

func createCamera(
	immutable entities.Immutable,
	viewport rectangles.Rectangle,
	fov float64,
	index uint,
) Camera {
	out := camera{
		immutable: immutable,
		viewport:  viewport,
		fov:       fov,
		index:     index,
	}

	return &out
}

// Hash returns the hash
func (obj *camera) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Viewport returns the viewport
func (obj *camera) Viewport() rectangles.Rectangle {
	return obj.viewport
}

// FieldOfView returns the field of view
func (obj *camera) FieldOfView() float64 {
	return obj.fov
}

// Index returns the index
func (obj *camera) Index() uint {
	return obj.index
}

// CreatedOn returns the creation time
func (obj *camera) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
