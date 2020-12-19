package nodes

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

type nodes struct {
	all []Node
}

func createNodes(
	all []Node,
) Nodes {
	out := nodes{
		all: all,
	}

	return &out
}

// All return all nodes
func (obj *nodes) All() []Node {
	return obj.all
}

// Render render the nodes
func (obj *nodes) Render(delta time.Duration, camera cameras.Camera, globalSpace spaces.Space, renderApp renders.Application) error {
	for _, oneNode := range obj.all {
		err := oneNode.Render(delta, camera, globalSpace, renderApp)
		if err != nil {
			return err
		}
	}

	return nil
}

// Camera returns the node and space of the camera at index
func (obj *nodes) Camera(index uint) (cameras.Camera, spaces.Space, error) {
	return obj.camera(index, nil)
}

func (obj *nodes) camera(index uint, prev spaces.Space) (cameras.Camera, spaces.Space, error) {
	for _, oneNode := range obj.all {
		if !oneNode.HasContent() {
			continue
		}

		nodeSpace := oneNode.Space()
		content := oneNode.Content()
		if !content.IsCamera() {
			continue
		}

		camera := content.Camera()
		if camera.Index() != index {
			continue
		}

		if prev == nil {
			return camera, nodeSpace, nil
		}

		worldSpace := nodeSpace.Add(prev)
		return camera, worldSpace, nil
	}

	for _, oneNode := range obj.all {
		if !oneNode.HasChildren() {
			continue
		}

		nodeSpace := oneNode.Space()
		children := oneNode.Children()
		for _, oneChildrenNode := range children {
			childrenNodeSpace := oneChildrenNode.Space()
			worldSpace := nodeSpace.Add(childrenNodeSpace)
			if prev != nil {
				worldSpace.Add(prev)
			}

			cam, space, err := obj.camera(index, worldSpace)
			if err != nil {
				continue
			}

			return cam, space, nil
		}
	}

	str := fmt.Sprintf("the camera (index: %d) could not be found in the nodes", index)
	return nil, nil, errors.New(str)
}
