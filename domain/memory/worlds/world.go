package worlds

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
)

type world struct {
	id                *uuid.UUID
	currentSceneIndex uint
	scenes            []scenes.Scene
}

func createWorld(
	id *uuid.UUID,
	currentSceneIndex uint,
	scenes []scenes.Scene,
) World {
	out := world{
		id:                id,
		currentSceneIndex: currentSceneIndex,
		scenes:            scenes,
	}

	return &out
}

// ID returns the id
func (obj *world) ID() *uuid.UUID {
	return obj.id
}

// CurrentSceneIndex returns the currentSceneIndex
func (obj *world) CurrentSceneIndex() uint {
	return obj.currentSceneIndex
}

// Scenes returns the scenes
func (obj *world) Scenes() []scenes.Scene {
	return obj.scenes
}
