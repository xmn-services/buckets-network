package models

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type model struct {
	immutable entities.Immutable
	geo       geometries.Geometry
	mat       materials.Material
	variable  string
}

func createModel(
	immutable entities.Immutable,
	geo geometries.Geometry,
	mat materials.Material,
	variable string,
) Model {
	out := model{
		immutable: immutable,
		geo:       geo,
		mat:       mat,
		variable:  variable,
	}

	return &out
}

// Hash returns the hash
func (obj *model) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Geometry returns the geometry
func (obj *model) Geometry() geometries.Geometry {
	return obj.geo
}

// Material returns the material
func (obj *model) Material() materials.Material {
	return obj.mat
}

// Variable returns the variable
func (obj *model) Variable() string {
	return obj.variable
}

// CreatedOn returns the creation time
func (obj *model) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
