package nodes

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the node builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithPosition(pos fl32.Vec3) Builder
	WithPositionVariable(posVar string) Builder
	WithOrientationAngle(angle float32) Builder
	WithOrientationDirection(direction fl32.Vec3) Builder
	WithOrientationVariable(oriVar string) Builder
	WithModel(model models.Model) Builder
	WithCamera(camera cameras.Camera) Builder
	WithChildren(children []Node) Builder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	ID() *uuid.UUID
	Position() Position
	Orientation() Orientation
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
}

// Content represents the node content
type Content interface {
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}

// Position represents a position
type Position interface {
	Vector() fl32.Vec3
	Variable() string
}

// Orientation represents an orientation
type Orientation interface {
	Angle() float32
	Direction() fl32.Vec3
	Variable() string
}
