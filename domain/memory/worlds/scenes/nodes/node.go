package nodes

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type node struct {
	immutable entities.Immutable
	space     Space
	content   Content
	nodes     []Node
}

func createNode(
	immutable entities.Immutable,
	space Space,
) Node {
	return createNodeInternally(immutable, space, nil, nil)
}

func createNodeWithContent(
	immutable entities.Immutable,
	space Space,
	content Content,
) Node {
	return createNodeInternally(immutable, space, content, nil)
}

func createNodeWithNodes(
	immutable entities.Immutable,
	space Space,
	nodes []Node,
) Node {
	return createNodeInternally(immutable, space, nil, nodes)
}

func createNodeWithContentAndNodes(
	immutable entities.Immutable,
	space Space,
	content Content,
	nodes []Node,
) Node {
	return createNodeInternally(immutable, space, content, nodes)
}

func createNodeInternally(
	immutable entities.Immutable,
	space Space,
	content Content,
	nodes []Node,
) Node {
	out := node{
		immutable: immutable,
		space:     space,
		content:   content,
		nodes:     nodes,
	}

	return &out
}

// Hash returns the hash
func (obj *node) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Space returns the space
func (obj *node) Space() Space {
	return obj.space
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
