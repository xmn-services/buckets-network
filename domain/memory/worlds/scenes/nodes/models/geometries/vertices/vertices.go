package vertices

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type vertices struct {
	mutable entities.Mutable
	list    []vertex.Vertex
	typ     Type
}

func createVertices(
	mutable entities.Mutable,
	list []vertex.Vertex,
	typ Type,
) Vertices {
	out := vertices{
		mutable: mutable,
		list:    list,
		typ:     typ,
	}

	return &out
}

// Hash returns the hash
func (obj *vertices) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of vertex
func (obj *vertices) All() []vertex.Vertex {
	return obj.list
}

// Type returns the type
func (obj *vertices) Type() Type {
	return obj.typ
}

// CreatedOn returns the creation time
func (obj *vertices) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *vertices) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}
