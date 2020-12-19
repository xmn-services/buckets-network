package nodes

import (
	"github.com/go-gl/mathgl/mgl32"
	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type node struct {
	original    domain_nodes.Node
	program     programs.Program
	pos         mgl32.Vec3
	orientation mgl32.Vec4
	content     Content
	children    []Node
}

func createNode(
	original domain_nodes.Node,
	program programs.Program,
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
) Node {
	return createNodeInternally(
		original,
		program,
		pos,
		orientation,
		nil,
		nil,
	)
}

func createNodeWithContent(
	original domain_nodes.Node,
	program programs.Program,
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
	content Content,
) Node {
	return createNodeInternally(
		original,
		program,
		pos,
		orientation,
		content,
		nil,
	)
}

func createNodeWithChildren(
	original domain_nodes.Node,
	program programs.Program,
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
	children []Node,
) Node {
	return createNodeInternally(
		original,
		program,
		pos,
		orientation,
		nil,
		children,
	)
}

func createNodeWithContentAndChildren(
	original domain_nodes.Node,
	program programs.Program,
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
	content Content,
	children []Node,
) Node {
	return createNodeInternally(
		original,
		program,
		pos,
		orientation,
		content,
		children,
	)
}

func createNodeInternally(
	original domain_nodes.Node,
	program programs.Program,
	pos mgl32.Vec3,
	orientation mgl32.Vec4,
	content Content,
	children []Node,
) Node {
	out := node{
		original:    original,
		program:     program,
		pos:         pos,
		orientation: orientation,
		content:     content,
		children:    children,
	}

	return &out
}

// Original returns the original node
func (obj *node) Original() domain_nodes.Node {
	return obj.original
}

// Program returns the program
func (obj *node) Program() programs.Program {
	return obj.program
}

// Position returns the position
func (obj *node) Position() mgl32.Vec3 {
	return obj.pos
}

// Orientation returns the orientation
func (obj *node) Orientation() mgl32.Vec4 {
	return obj.orientation
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
func (obj *node) Render(cameraIndex uint) error {
	if obj.HasContent() {
		if obj.content.IsCamera() {
			camera := obj.content.Camera()
			if camera.Original().Index() == cameraIndex {
				err := obj.content.Camera().Render()
				if err != nil {
					return err
				}
			}
		}

		if obj.content.IsModel() {
			err := obj.content.Model().Render(
				obj.Position(),
				obj.Orientation(),
			)

			if err != nil {
				return err
			}
		}
	}

	if obj.HasChildren() {
		for _, oneChildren := range obj.children {
			err := oneChildren.Render(cameraIndex)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
