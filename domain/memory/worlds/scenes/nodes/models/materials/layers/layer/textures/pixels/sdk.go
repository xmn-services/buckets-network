package pixels

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"
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

// Factory represents a pixels factory
type Factory interface {
	Create() (Pixels, error)
}

// Builder represents the pixels builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithoutHash() Builder
	WithPixels(pixels []pixel.Pixel) Builder
	CreatedOn(createdOn time.Time) Builder
	LastUpdatedOn(lastUpdatedOn time.Time) Builder
	Now() (Pixels, error)
}

// Pixels represents pixels
type Pixels interface {
	entities.Immutable
	All() []pixel.Pixel
	Amount() uint
}
