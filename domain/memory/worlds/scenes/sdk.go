package scenes

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// CurrentSceneIndex represents the current scene index
const CurrentSceneIndex = 0

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder, CurrentSceneIndex)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Factory represents a scene factory
type Factory interface {
	Create() (Scene, error)
}

// Builder represents the scene builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithNodes(nodes []nodes.Node) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	entities.Immutable
	Index() uint
	Add(node nodes.Node) error
	Camera(index uint) (cameras.Camera, error)
	HasNodes() bool
	Nodes() []nodes.Node
}
