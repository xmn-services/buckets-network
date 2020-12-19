package shaders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/shaders/shader"
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
	mutableBuilder := entities.NewMutableBuilder()
	return createBuilder(hashAdapter, mutableBuilder)
}

// Factory represents a shaders factory
type Factory interface {
	Create() (Shaders, error)
}

// Builder represents the shaders builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithShaders(shaders []shader.Shader) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	entities.Mutable
	All() []shader.Shader
	IsEmpty() bool
	IsVertex() bool
	IsFragment() bool
}
