package scenes

import (
	domain_scenes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
)

// NewBuilder creates a new builder instance
func NewBuilder(defaultCameraIndex uint) Builder {
	nodeBuilder := nodes.NewBuilder()
	return createBuilder(nodeBuilder, defaultCameraIndex)
}

// Builder represents the scene builder
type Builder interface {
	Create() Builder
	WithScene(scene domain_scenes.Scene) Builder
	WithCameraIndex(cameraIndex uint) Builder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	Original() domain_scenes.Scene
	CameraIndex() uint
	HasNodes() bool
	Nodes() []nodes.Node
	Render() error
}
