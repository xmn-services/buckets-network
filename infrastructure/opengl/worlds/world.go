package worlds

import (
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	domain_worlds "github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/scenes"
)

type world struct {
	original   domain_worlds.World
	sceneIndex uint
	scene      scenes.Scene
}

func createWorld(
	original domain_worlds.World,
	sceneIndex uint,
) World {
	return createWorldInternally(original, sceneIndex, nil)
}

func createWorldWithScenes(
	original domain_worlds.World,
	sceneIndex uint,
	scene scenes.Scene,
) World {
	return createWorldInternally(original, sceneIndex, scene)
}

func createWorldInternally(
	original domain_worlds.World,
	sceneIndex uint,
	scene scenes.Scene,
) World {
	out := world{
		original:   original,
		sceneIndex: sceneIndex,
		scene:      scene,
	}

	return &out
}

// Original returns the original
func (obj *world) Original() domain_worlds.World {
	return obj.original
}

// SceneIndex returns the scene index
func (obj *world) SceneIndex() uint {
	return obj.sceneIndex
}

// HasScene returns true if there is a scene, false otherwise
func (obj *world) HasScene() bool {
	return obj.scene != nil
}

// Scene returns the scene, if any
func (obj *world) Scene() scenes.Scene {
	return obj.scene
}

// Render renders the world
func (obj *world) Render(delta time.Duration) error {
	gl.ClearDepth(10.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	if !obj.HasScene() {
		return nil
	}

	return obj.scene.Render(delta)
}
