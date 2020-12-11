package program

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type program struct {
	scene      hash.Hash
	materials  materials.Materials
	identifier uint32
}

func createProgram(
	scene hash.Hash,
	materials materials.Materials,
	identifier uint32,
) Program {
	out := program{
		scene:      scene,
		materials:  materials,
		identifier: identifier,
	}

	return &out
}

// Scene returns the scene hash
func (obj *program) Scene() hash.Hash {
	return obj.scene
}

// Compiled returns the compiled material
func (obj *program) Compiled() materials.Materials {
	return obj.materials
}

// Identifier returns the identifier
func (obj *program) Identifier() uint32 {
	return obj.identifier
}
