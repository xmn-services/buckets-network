package scenes

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the scene builder
type Builder interface {
	Create() Builder
	WithNodes(nodes []nodes.Node) Builder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	entities.Immutable
	Nodes() []nodes.Node
}
