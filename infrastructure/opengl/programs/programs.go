package programs

import (
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type programs struct {
	list []program.Program
	mp   map[string]program.Program
}

func createPrograms(
	list []program.Program,
	mp map[string]program.Program,
) Programs {
	out := programs{
		list: list,
		mp:   mp,
	}

	return &out
}

// All return the programs
func (obj *programs) All() []program.Program {
	return obj.list
}

// Fetch fetches a program by scene hash
func (obj *programs) Fetch(scene hash.Hash) (program.Program, error) {
	keyname := scene.String()
	if prog, ok := obj.mp[keyname]; ok {
		return prog, nil
	}

	str := fmt.Sprintf("there is no program for scene (hash: %s)", keyname)
	return nil, errors.New(str)
}
