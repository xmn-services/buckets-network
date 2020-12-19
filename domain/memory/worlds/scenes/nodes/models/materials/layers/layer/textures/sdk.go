package textures

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents a texture builder
type Builder interface {
	Create() Builder
	WithDimension(dimension ints.Rectangle) Builder
	WithPixels(pixels rows.Rows) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Texture, error)
}

// Texture represents a texture
type Texture interface {
	entities.Immutable
	Dimension() ints.Rectangle
	Pixels() rows.Rows
}

// Repository represents the texture repository
type Repository interface {
	Retrieve(path string) (Texture, error)
}

// Service represents a texture service
type Service interface {
	Save(tex Texture) error
	SaveAll(tex []Texture) error
}
