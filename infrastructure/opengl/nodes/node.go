package nodes

import (
	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

type node struct {
	original domain_nodes.Node
	space    spaces.Space
	content  Content
	children []Node
}

func createNode(
	original domain_nodes.Node,
	space spaces.Space,
) Node {
	return createNodeInternally(
		original,
		space,
		nil,
		nil,
	)
}

func createNodeWithContent(
	original domain_nodes.Node,
	space spaces.Space,
	content Content,
) Node {
	return createNodeInternally(
		original,
		space,
		content,
		nil,
	)
}

func createNodeWithChildren(
	original domain_nodes.Node,
	space spaces.Space,
	children []Node,
) Node {
	return createNodeInternally(
		original,
		space,
		nil,
		children,
	)
}

func createNodeWithContentAndChildren(
	original domain_nodes.Node,
	space spaces.Space,
	content Content,
	children []Node,
) Node {
	return createNodeInternally(
		original,
		space,
		content,
		children,
	)
}

func createNodeInternally(
	original domain_nodes.Node,
	space spaces.Space,
	content Content,
	children []Node,
) Node {
	out := node{
		original: original,
		space:    space,
		content:  content,
		children: children,
	}

	return &out
}

// Original returns the original node
func (obj *node) Original() domain_nodes.Node {
	return obj.original
}

// Space returns the space
func (obj *node) Space() spaces.Space {
	return obj.space
}

// HasContent returns true if there is content, false otherwise
func (obj *node) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *node) Content() Content {
	return obj.content
}

// HasChildren returns true if there is children, false otherwise
func (obj *node) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *node) Children() []Node {
	return obj.children
}

// Render renders the node
func (obj *node) Render(camera cameras.Camera, globalSpace spaces.Space) error {
	if obj.HasContent() {
		if !obj.content.IsModel() {
			return nil
		}

		err := obj.content.Model().Render(
			camera,
			obj.Space(),
		)

		if err != nil {
			return err
		}
	}

	if obj.HasChildren() {
		for _, oneChildren := range obj.children {
			err := oneChildren.Render(camera, globalSpace)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
