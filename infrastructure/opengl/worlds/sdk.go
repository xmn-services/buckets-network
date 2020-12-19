package worlds

import (
	domain_worlds "github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/scenes"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	defaultSceneIndex uint,
	defaultCameraIndex uint,
) Builder {
	sceneBuilder := scenes.NewBuilder(defaultCameraIndex)
	return createBuilder(sceneBuilder, defaultSceneIndex)
}

// Builder represents a world builder
type Builder interface {
	Create() Builder
	WithWorld(world domain_worlds.World) Builder
	WithSceneIndex(sceneIndex uint) Builder
	Now() (World, error)
}

// World represents a world
type World interface {
	Original() domain_worlds.World
	SceneIndex() uint
	HasScene() bool
	Scene() scenes.Scene
	Render() error
}
