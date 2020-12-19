package nodes

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type node struct {
	immutable   entities.Immutable
	position    fl32.Vec3
	orientation Orientation
	content     Content
	nodes       []Node
}

func createNode(
	immutable entities.Immutable,
	position fl32.Vec3,
	orientation Orientation,
) Node {
	return createNodeInternally(immutable, position, orientation, nil, nil)
}

func createNodeWithContent(
	immutable entities.Immutable,
	position fl32.Vec3,
	orientation Orientation,
	content Content,
) Node {
	return createNodeInternally(immutable, position, orientation, content, nil)
}

func createNodeWithNodes(
	immutable entities.Immutable,
	position fl32.Vec3,
	orientation Orientation,
	nodes []Node,
) Node {
	return createNodeInternally(immutable, position, orientation, nil, nodes)
}

func createNodeWithContentAndNodes(
	immutable entities.Immutable,
	position fl32.Vec3,
	orientation Orientation,
	content Content,
	nodes []Node,
) Node {
	return createNodeInternally(immutable, position, orientation, content, nodes)
}

func createNodeInternally(
	immutable entities.Immutable,
	position fl32.Vec3,
	orientation Orientation,
	content Content,
	nodes []Node,
) Node {
	out := node{
		immutable:   immutable,
		position:    position,
		orientation: orientation,
		content:     content,
		nodes:       nodes,
	}

	return &out
}

// Hash returns the hash
func (obj *node) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Camera returns the camera at index, if any
func (obj *node) Camera(index uint) (cameras.Camera, error) {
	if obj.HasContent() {
		content := obj.Content()
		if content.IsCamera() {
			camera := content.Camera()
			if camera.Index() == index {
				return camera, nil
			}
		}
	}

	if obj.HasChildren() {
		children := obj.Children()
		for _, oneNode := range children {
			cam, err := oneNode.Camera(index)
			if err != nil {
				continue
			}

			return cam, nil
		}
	}

	str := fmt.Sprintf("the camera (index: %d) could not be found in the node (hash: %s)", index, obj.Hash().String())
	return nil, errors.New(str)
}

// Position returns the position
func (obj *node) Position() fl32.Vec3 {
	return obj.position
}

// Orientation returns the orientation
func (obj *node) Orientation() Orientation {
	return obj.orientation
}

// CreatedOn returns the creation time
func (obj *node) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasContent returns true if there is content, false otherwise
func (obj *node) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *node) Content() Content {
	return obj.content
}

// HasChildren returns true if there is children nodes, false otherwise
func (obj *node) HasChildren() bool {
	return obj.nodes != nil
}

// Children returns the children nodes, if any
func (obj *node) Children() []Node {
	return obj.nodes
}
