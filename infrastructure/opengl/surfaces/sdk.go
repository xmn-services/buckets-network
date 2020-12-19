package surfaces

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces/surface"
)

// NewBuilder creates a new surface builder
func NewBuilder() Builder {
	surfaceBuilder := surface.NewBuilder()
	return createBuilder(surfaceBuilder)
}

// Builder represents a surfaces builder
type Builder interface {
	Create() Builder
	WithProgram(prog programs.Program) Builder
	WithRenders(renders renders.Renders) Builder
	Now() (Surfaces, error)
}

// Surfaces represents surfaces
type Surfaces interface {
	All() []surface.Surface
}
