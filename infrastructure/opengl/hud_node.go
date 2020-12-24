package opengl

import uuid "github.com/satori/go.uuid"

type hudNode struct {
	id          *uuid.UUID
	pos         HudPosition
	orientation Orientation
	content     HudNodeContent
	children    []HudNode
}

func createHudNode(
	id *uuid.UUID,
	pos HudPosition,
	orientation Orientation,
) HudNode {
	return createHudNodeInternally(id, pos, orientation, nil, nil)
}

func createHudNodeWithContent(
	id *uuid.UUID,
	pos HudPosition,
	orientation Orientation,
	content HudNodeContent,
) HudNode {
	return createHudNodeInternally(id, pos, orientation, content, nil)
}

func createHudNodeWithChildren(
	id *uuid.UUID,
	pos HudPosition,
	orientation Orientation,
	children []HudNode,
) HudNode {
	return createHudNodeInternally(id, pos, orientation, nil, children)
}

func createHudNodeWithContentAndChildren(
	id *uuid.UUID,
	pos HudPosition,
	orientation Orientation,
	content HudNodeContent,
	children []HudNode,
) HudNode {
	return createHudNodeInternally(id, pos, orientation, content, children)
}

func createHudNodeInternally(
	id *uuid.UUID,
	pos HudPosition,
	orientation Orientation,
	content HudNodeContent,
	children []HudNode,
) HudNode {
	out := hudNode{
		id:          id,
		pos:         pos,
		orientation: orientation,
		content:     content,
		children:    children,
	}

	return &out
}

// ID returns the id
func (obj *hudNode) ID() *uuid.UUID {
	return obj.id
}

// Position returns the position
func (obj *hudNode) Position() HudPosition {
	return obj.pos
}

// Orientation returns the orientation
func (obj *hudNode) Orientation() Orientation {
	return obj.orientation
}

// HasContent returns true if there is content, false otherwise
func (obj *hudNode) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *hudNode) Content() HudNodeContent {
	return obj.content
}

// HasChildren returns true if there is children, false otherwise
func (obj *hudNode) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *hudNode) Children() []HudNode {
	return obj.children
}
