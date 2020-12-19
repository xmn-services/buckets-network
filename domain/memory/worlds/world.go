package worlds

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type world struct {
	immutable entities.Immutable
	list      []scenes.Scene
	mp        map[string]scenes.Scene
}

func createWorld(
	immutable entities.Immutable,
) World {
	return createWorldInternally(immutable, nil, nil)
}

func createWorldWithScene(
	immutable entities.Immutable,
	list []scenes.Scene,
	mp map[string]scenes.Scene,
) World {
	return createWorldInternally(immutable, list, mp)
}

func createWorldInternally(
	immutable entities.Immutable,
	list []scenes.Scene,
	mp map[string]scenes.Scene,
) World {
	out := world{
		immutable: immutable,
		list:      list,
		mp:        mp,
	}

	return &out
}

// Hash returns the hash
func (obj *world) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Add adds a scene instance
func (obj *world) Add(scene scenes.Scene) error {
	if obj.mp == nil {
		obj.mp = map[string]scenes.Scene{}
	}

	keyname := scene.Hash().String()
	if existingScene, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the scene (hash: %s) already exists", existingScene.Hash().String())
		return errors.New(str)
	}

	obj.mp[keyname] = scene
	obj.list = append(obj.list, scene)
	return nil
}

// Scene returns the scene at index, if any
func (obj *world) Scene(index uint) (scenes.Scene, error) {
	for _, oneScene := range obj.list {
		if oneScene.Index() != index {
			continue
		}

		return oneScene, nil
	}

	str := fmt.Sprintf("the scene (index: %d) could not be found in the world (hash: %s)", index, obj.Hash().String())
	return nil, errors.New(str)
}

// HasScenes returns true if there is scenes, false otherwise
func (obj *world) HasScenes() bool {
	return obj.list != nil
}

// Scenes returns the scenes
func (obj *world) Scenes() []scenes.Scene {
	return obj.list
}

// CreatedOn returns the creation time
func (obj *world) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
