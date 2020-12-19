package nodes

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents the node builder
type Builder interface {
	Create() Builder
	WithPosition(pos fl32.Vec3) Builder
	WithOrientationAngle(angle float32) Builder
	WithOrientationDirection(direction fl32.Vec3) Builder
	WithModel(model models.Model) Builder
	WithCamera(camera cameras.Camera) Builder
	WithChildren(children []Node) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	entities.Immutable
	Camera(index uint) (cameras.Camera, error)
	Position() fl32.Vec3
	Orientation() Orientation
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
}

// Orientation represents an orientation
type Orientation interface {
	Angle() float32
	Direction() fl32.Vec3
}

// Content represents the node content
type Content interface {
	Hash() hash.Hash
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}
