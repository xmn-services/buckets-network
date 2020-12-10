package scenes

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type scene struct {
	immutable entities.Immutable
	nodes     []nodes.Node
}

func createScene(
	immutable entities.Immutable,
) Scene {
	return createSceneInternally(immutable, nil)
}

func createSceneWithNodes(
	immutable entities.Immutable,
	nodes []nodes.Node,
) Scene {
	return createSceneInternally(immutable, nodes)
}

func createSceneInternally(
	immutable entities.Immutable,
	nodes []nodes.Node,
) Scene {
	out := scene{
		immutable: immutable,
		nodes:     nodes,
	}

	return &out
}

// Hash returns the hash
func (obj *scene) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// CreatedOn returns the creation time
func (obj *scene) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *scene) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes, if any
func (obj *scene) Nodes() []nodes.Node {
	return obj.nodes
}
