package worlds

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents a world builder
type Builder interface {
	Create() Builder
	WithScenes(scenes []scenes.Scene) Builder
	Now() (World, error)
}

// World represents a world
type World interface {
	entities.Immutable
	Scenes() []scenes.Scene
}
