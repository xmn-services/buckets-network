package nodes

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math"
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
	WithPosition(pos math.Vec3) Builder
	WithRight(right math.Vec3) Builder
	WithUp(up math.Vec3) Builder
	WithModel(model models.Model) Builder
	WithCamera(camera cameras.Camera) Builder
	WithChildren(children []Node) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	entities.Immutable
	Space() Space
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
}

// Space represents the position and orientation of the node
type Space interface {
	Position() math.Vec3
	Right() math.Vec3
	Up() math.Vec3
}

// Content represents the node content
type Content interface {
	Hash() hash.Hash
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}
