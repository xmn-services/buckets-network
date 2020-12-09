package shaders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the shaders builder
type Builder interface {
	Create() Builder
	WithShaders(shaders []shader.Shader) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Shaders, error)
}

// Shaders represents shaders
type Shaders interface {
	entities.Immutable
	All() []shader.Shader
	IsVertex() bool
	IsFragment() bool
}
