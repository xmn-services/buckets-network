package worlds

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	sceneFactory := scenes.NewFactory()
	return createBuilder(hashAdapter, immutableBuilder, sceneFactory)
}

// Factory represents a world factory
type Factory interface {
	Create() (World, error)
}

// Builder represents a world builder
type Builder interface {
	Create() Builder
	WithScenes(scenes []scenes.Scene) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (World, error)
}

// World represents a world
type World interface {
	entities.Immutable
	Scenes() []scenes.Scene
}
