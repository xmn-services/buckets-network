package models

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the model builder
type Builder interface {
	Create() Builder
	WithGeometry(geo geometries.Geometry) Builder
	WithMaterial(material materials.Material) Builder
	Now() (Model, error)
}

// Model represents a model
type Model interface {
	entities.Immutable
	Geometry() geometries.Geometry
	Material() materials.Material
}
