package worlds

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a world builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithCurrentSceneIndex(currentSceneIndex uint) Builder
	WithScenes(scenes []scenes.Scene) Builder
	Now() (World, error)
}

// World represents a world
type World interface {
	ID() *uuid.UUID
	CurrentSceneIndex() uint
	Scenes() []scenes.Scene
}
