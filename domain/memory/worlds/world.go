package worlds

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type world struct {
	immutable entities.Immutable
	scenes    []scenes.Scene
}

func createWorld(
	immutable entities.Immutable,
	scenes []scenes.Scene,
) World {
	out := world{
		immutable: immutable,
		scenes:    scenes,
	}

	return &out
}

// Hash returns the hash
func (obj *world) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Scenes returns the scenes
func (obj *world) Scenes() []scenes.Scene {
	return obj.scenes
}

// CreatedOn returns the creation time
func (obj *world) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
