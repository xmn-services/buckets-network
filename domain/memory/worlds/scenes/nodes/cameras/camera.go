package cameras

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type camera struct {
	immutable  entities.Immutable
	index      uint
	projection Projection
	lookAt     LookAt
}

func createCamera(
	immutable entities.Immutable,
	index uint,
	projection Projection,
	lookAt LookAt,
) Camera {
	out := camera{
		immutable:  immutable,
		index:      index,
		projection: projection,
		lookAt:     lookAt,
	}

	return &out
}

// Hash returns the hash
func (obj *camera) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Index returns the camera index
func (obj *camera) Index() uint {
	return obj.index
}

// Projection returns the camera projection
func (obj *camera) Projection() Projection {
	return obj.projection
}

// LookAt returns the camera lookAt
func (obj *camera) LookAt() LookAt {
	return obj.lookAt
}

// CreatedOn returns the creation time
func (obj *camera) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
