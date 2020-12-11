package materials

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials/material"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders/shader"
)

type materials struct {
	list []material.Material
}

func createMaterials(
	list []material.Material,
) Materials {
	out := materials{
		list: list,
	}

	return &out
}

// All return all compiled materials
func (obj *materials) All() []material.Material {
	return obj.list
}

// CompiledShaders returns all compiled shaders of the materials
func (obj *materials) CompiledShaders() []shader.Shader {
	out := []shader.Shader{}
	for _, oneMat := range obj.list {
		sub := oneMat.Compiled().CompiledShaders()
		out = append(out, sub...)
	}

	return out
}
