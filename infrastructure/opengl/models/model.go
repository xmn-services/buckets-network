package models

import (
	"math"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type model struct {
	model           models.Model
	typ             Type
	vao             uint32
	vertexAmount    int32
	uniformVariable int32
	prog            programs.Program
	material        materials.Material
}

func createModel(
	domainModel models.Model,
	typ Type,
	vao uint32,
	vertexAmount int32,
	uniformVariable int32,
	prog programs.Program,
	material materials.Material,
) Model {
	out := model{
		model:           domainModel,
		typ:             typ,
		vao:             vao,
		vertexAmount:    vertexAmount,
		uniformVariable: uniformVariable,
		prog:            prog,
		material:        material,
	}

	return &out
}

// Model returns the domain model
func (obj *model) Model() models.Model {
	return obj.model
}

// Type returns the type
func (obj *model) Type() Type {
	return obj.typ
}

// VAO returns the vao
func (obj *model) VAO() uint32 {
	return obj.vao
}

// VertexAmount returns the vertex amountRender(pos mgl32.Vec3, orientation mgl32.Vec4) error
func (obj *model) VertexAmount() int32 {
	return obj.vertexAmount
}

// UniformVariable returns the uniform variable
func (obj *model) UniformVariable() int32 {
	return obj.uniformVariable
}

// Program returns the program
func (obj *model) Program() programs.Program {
	return obj.prog
}

// Material returns the material
func (obj *model) Material() materials.Material {
	return obj.material
}

// Render renders the model
func (obj *model) Render(pos mgl32.Vec3, orientation mgl32.Vec4) error {
	// use the program:
	identifier := obj.Program().Identifier()
	gl.UseProgram(identifier)

	// rotate then translate:
	radAngle := float32(orientation[3] * math.Pi / 180)
	rorateMat := mgl32.HomogRotate3D(radAngle, mgl32.Vec3{orientation[0], orientation[1], orientation[2]})

	// translate:
	transMat := mgl32.Translate3D(pos[0], pos[1], pos[2])

	// model matrix:
	modelMat := transMat.Add(rorateMat)

	// apply the model matrix:
	uniform := obj.UniformVariable()
	gl.UniformMatrix4fv(uniform, 1, false, &modelMat[0])

	// vao:
	vao := obj.VAO()
	gl.BindVertexArray(vao)

	// material:
	err := obj.Material().Render()
	if err != nil {
		return err
	}

	// draw:
	amount := obj.VertexAmount()
	gl.DrawArrays(gl.TRIANGLES, 0, amount)
	return nil
}
