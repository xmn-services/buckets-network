package programs

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials/material"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	programBuilder := program.NewBuilder()
	programsBuilder := NewBuilder()
	materialBuilder := material.NewBuilder()
	materialsBuilder := materials.NewBuilder()
	layerBuilder := layer.NewBuilder()
	layersBuilder := layers.NewBuilder()
	shadersApplication := shaders.NewApplication()
	return createApplication(
		programBuilder,
		programsBuilder,
		materialBuilder,
		materialsBuilder,
		layerBuilder,
		layersBuilder,
		shadersApplication,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Application represents a program application
type Application interface {
	Execute(world worlds.World) (Programs, error)
}

// Builder represents programs builder
type Builder interface {
	Create() Builder
	WithPrograms(progs []program.Program) Builder
	Now() (Programs, error)
}

// Programs represents compiled programs
type Programs interface {
	All() []program.Program
}
