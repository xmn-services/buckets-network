package geometries

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type geometry struct {
	immutable entities.Immutable
	shaders   shaders.Shaders
	vertices  vertices.Vertices
}

func createGeometry(
	immutable entities.Immutable,
	shaders shaders.Shaders,
	vertices vertices.Vertices,
) Geometry {
	out := geometry{
		immutable: immutable,
		shaders:   shaders,
		vertices:  vertices,
	}

	return &out
}

// Hash returns the hash
func (obj *geometry) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Shaders returns the shaders
func (obj *geometry) Shaders() shaders.Shaders {
	return obj.shaders
}

// Vertices returns the vertices
func (obj *geometry) Vertices() vertices.Vertices {
	return obj.vertices
}

// CreatedOn returns the creation time
func (obj *geometry) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
