package program

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithScene(scene hash.Hash) Builder
	WithCompiledMaterials(materials materials.Materials) Builder
	WithIdentifier(identifier uint32) Builder
	Now() (Program, error)
}

// Program represents a compiled program
type Program interface {
	Scene() hash.Hash
	Compiled() materials.Materials
	Identifier() uint32
}
