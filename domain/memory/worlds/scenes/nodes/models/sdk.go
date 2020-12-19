package models

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents the model builder
type Builder interface {
	Create() Builder
	WithGeometry(geo geometries.Geometry) Builder
	WithMaterial(material materials.Material) Builder
	WithVariable(variable string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Model, error)
}

// Model represents a model
type Model interface {
	entities.Immutable
	Geometry() geometries.Geometry
	Material() materials.Material
	Variable() string
}
