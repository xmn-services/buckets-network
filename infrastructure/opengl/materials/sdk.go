package materials

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	domain_materials "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	layersBuilder := layers.NewBuilder()
	return createBuilder(layersBuilder)
}

// Builder represents a material builder
type Builder interface {
	Create() Builder
	WithMaterial(material domain_materials.Material) Builder
	WithProgram(prog programs.Program) Builder
	Now() (Material, error)
}

// Material represents a material
type Material interface {
	Original() domain_materials.Material
	Alpha() uint8
	Viewport() ints.Rectangle
	Layers() layers.Layers
}
