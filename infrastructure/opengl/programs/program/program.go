package program

import "github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"

type program struct {
	materials  materials.Materials
	identifier uint32
}

func createProgram(
	materials materials.Materials,
	identifier uint32,
) Program {
	out := program{
		materials:  materials,
		identifier: identifier,
	}

	return &out
}

// Compiled returns the compiled material
func (obj *program) Compiled() materials.Materials {
	return obj.materials
}

// Identifier returns the identifier
func (obj *program) Identifier() uint32 {
	return obj.identifier
}
