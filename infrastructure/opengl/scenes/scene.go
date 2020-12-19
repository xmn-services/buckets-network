package scenes

import (
	domain_scenes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
)

type scene struct {
	original    domain_scenes.Scene
	cameraIndex uint
	nodes       []nodes.Node
}

func createScene(
	original domain_scenes.Scene,
	cameraIndex uint,
) Scene {
	return createSceneInternally(original, cameraIndex, nil)
}

func createSceneWithNodes(
	original domain_scenes.Scene,
	cameraIndex uint,
	nodes []nodes.Node,
) Scene {
	return createSceneInternally(original, cameraIndex, nodes)
}

func createSceneInternally(
	original domain_scenes.Scene,
	cameraIndex uint,
	nodes []nodes.Node,
) Scene {
	out := scene{
		original:    original,
		cameraIndex: cameraIndex,
		nodes:       nodes,
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
func (obj *scene) Nodes() []nodes.Node {
	return obj.nodes
}

// Render renders the scene
func (obj *scene) Render() error {
	if !obj.HasNodes() {
		return nil
	}

	for _, oneNode := range obj.nodes {
		err := oneNode.Render(obj.cameraIndex)
		if err != nil {
			return err
		}
	}

	return nil
}
