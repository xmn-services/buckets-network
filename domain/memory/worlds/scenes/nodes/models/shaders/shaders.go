package shaders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type shaders struct {
	mutable entities.Mutable
	list    []shader.Shader
}

func createShaders(
	mutable entities.Mutable,
	list []shader.Shader,
) Shaders {
	out := shaders{
		mutable: mutable,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *shaders) Hash() hash.Hash {
	return obj.mutable.Hash()
}

// All returns the list of shader
func (obj *shaders) All() []shader.Shader {
	return obj.list
}

// CreatedOn returns the creation time
func (obj *shaders) CreatedOn() time.Time {
	return obj.mutable.CreatedOn()
}

// LastUpdatedOn returns the lasUpdatedOn time
func (obj *shaders) LastUpdatedOn() time.Time {
	return obj.mutable.LastUpdatedOn()
}

// IsEmpty returns true if the list is empty, false otherwise
func (obj *shaders) IsEmpty() bool {
	return len(obj.list) <= 0
}

// IsVertex returns true if the shaders are vertex shaders, false otherwise
func (obj *shaders) IsVertex() bool {
	if obj.IsEmpty() {
		return false
	}

	return obj.list[0].Type().IsVertex()
}

// IsFragment returns true if the shaders are fragment shaders, false otherwise
func (obj *shaders) IsFragment() bool {
	if obj.IsEmpty() {
		return false
	}

	return obj.list[0].Type().IsFragment()
}
