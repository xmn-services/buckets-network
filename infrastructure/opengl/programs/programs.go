package programs

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
)

type programs struct {
	list []program.Program
}

func createPrograms(
	list []program.Program,
) Programs {
	out := programs{
		list: list,
	}

	return &out
}

// All return the programs
func (obj *programs) All() []program.Program {
	return obj.list
}
