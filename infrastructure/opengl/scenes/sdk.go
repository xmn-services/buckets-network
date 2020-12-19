package scenes

import (
	"time"

	domain_scenes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders/applications"
)

// NewBuilder creates a new builder instance
func NewBuilder(defaultCameraIndex uint) Builder {
	renderAppBuilder := applications.NewBuilder()
	nodesBuilder := nodes.NewBuilder()
	return createBuilder(renderAppBuilder, nodesBuilder, defaultCameraIndex)
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
	Nodes() nodes.Nodes
	Render(delta time.Duration) error
}
