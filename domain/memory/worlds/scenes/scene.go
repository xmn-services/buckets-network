package scenes

import (
	"errors"
	"fmt"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type scene struct {
	immutable entities.Immutable
	index     uint
	list      []nodes.Node
	mp        map[string]nodes.Node
}

func createScene(
	immutable entities.Immutable,
	index uint,
) Scene {
	return createSceneInternally(immutable, index, nil, nil)
}

func createSceneWithNodes(
	immutable entities.Immutable,
	index uint,
	list []nodes.Node,
	mp map[string]nodes.Node,
) Scene {
	return createSceneInternally(immutable, index, list, mp)
}

func createSceneInternally(
	immutable entities.Immutable,
	index uint,
	list []nodes.Node,
	mp map[string]nodes.Node,
) Scene {
	out := scene{
		immutable: immutable,
		index:     index,
		list:      list,
		mp:        mp,
	}

	return &out
}

// Hash returns the hash
func (obj *scene) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Index returns the index
func (obj *scene) Index() uint {
	return obj.index
}

// Camera returns the camera at index 0
func (obj *scene) Camera(index uint) (cameras.Camera, error) {
	for _, oneNode := range obj.list {
		cam, err := oneNode.Camera(index)
		if err != nil {
			continue
		}

		if cam == nil {
			continue
		}

		return cam, nil
	}

	str := fmt.Sprintf("the camera (index: %d) could not be found in the scene (hash: %s)", index, obj.Hash().String())
	return nil, errors.New(str)
}

// Add adds a node instance
func (obj *scene) Add(node nodes.Node) error {
	if obj.mp == nil {
		obj.mp = map[string]nodes.Node{}
	}

	keyname := node.Hash().String()
	if existingNode, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the node (hash: %s) already exists", existingNode.Hash().String())
		return errors.New(str)
	}

	obj.mp[keyname] = node
	obj.list = append(obj.list, node)
	return nil
}

// CreatedOn returns the creation time
func (obj *scene) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *scene) HasNodes() bool {
	return obj.list != nil
}

// Nodes returns the nodes, if any
func (obj *scene) Nodes() []nodes.Node {
	return obj.list
}
