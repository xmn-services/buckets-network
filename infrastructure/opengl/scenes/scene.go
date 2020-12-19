package scenes

import (
	"time"

	domain_scenes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders/applications"
)

type scene struct {
	renderAppBuilder applications.Builder
	original         domain_scenes.Scene
	cameraIndex      uint
	nodes            nodes.Nodes
}

func createScene(
	renderAppBuilder applications.Builder,
	original domain_scenes.Scene,
	cameraIndex uint,
) Scene {
	return createSceneInternally(renderAppBuilder, original, cameraIndex, nil)
}

func createSceneWithNodes(
	renderAppBuilder applications.Builder,
	original domain_scenes.Scene,
	cameraIndex uint,
	nodes nodes.Nodes,
) Scene {
	return createSceneInternally(renderAppBuilder, original, cameraIndex, nodes)
}

func createSceneInternally(
	renderAppBuilder applications.Builder,
	original domain_scenes.Scene,
	cameraIndex uint,
	nodes nodes.Nodes,
) Scene {
	out := scene{
		renderAppBuilder: renderAppBuilder,
		original:         original,
		cameraIndex:      cameraIndex,
		nodes:            nodes,
	}

	return &out
}

// Original returns the original scene
func (obj *scene) Original() domain_scenes.Scene {
	return obj.original
}

// CameraIndex returns the camera index
func (obj *scene) CameraIndex() uint {
	return obj.cameraIndex
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *scene) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes
func (obj *scene) Nodes() nodes.Nodes {
	return obj.nodes
}

// Render renders the scene
func (obj *scene) Render(delta time.Duration) error {
	if !obj.HasNodes() {
		return nil
	}

	// render all cameras contained in a surface, to a texture:
	renderApp, err := obj.renderAppBuilder.Create().WithNodes(obj.nodes).Now()
	if err != nil {
		return err
	}

	// find the camera:
	currentCamera, currentCameraSpace, err := obj.nodes.Camera(obj.cameraIndex)
	if err != nil {
		return err
	}

	// render the nodes:
	return obj.nodes.Render(delta, currentCamera, currentCameraSpace, renderApp)
}
