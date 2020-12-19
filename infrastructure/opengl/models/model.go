package models

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

type model struct {
	model           models.Model
	typ             Type
	vao             uint32
	vertexAmount    int32
	uniformVariable int32
	prog            programs.Program
	material        materials.Material
	angle           float32
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
		angle:           0.0,
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
func (obj *model) Render(delta time.Duration, camera cameras.Camera, space spaces.Space, renderApp renders.Application) error {
	//pos := space.Position()
	orientation := space.Orientation()

	identifier := obj.Program().Identifier()
	gl.UseProgram(identifier)

	// projection:
	projection := camera.Projection()
	projVariable := fmt.Sprintf(glStrPattern, projection.Variable())
	fov := projection.FieldOfView()
	aspectRatio := projection.AspectRation()
	near := projection.Near()
	far := projection.Far()

	projMat := mgl32.Perspective(
		mgl32.DegToRad(fov),
		aspectRatio,
		near,
		far,
	)

	projectionUniform := gl.GetUniformLocation(
		obj.prog.Identifier(),
		gl.Str(projVariable),
	)

	gl.UniformMatrix4fv(
		projectionUniform,
		1,
		false,
		&projMat[0],
	)

	// lookAt:
	lookAt := camera.LookAt()
	lookAtVariable := fmt.Sprintf(glStrPattern, lookAt.Variable())
	eye := lookAt.Eye() // Configure global settings
	center := lookAt.Center()
	up := lookAt.Up()

	lookAtMat := mgl32.LookAtV(
		mgl32.Vec3{eye[0], eye[1], eye[2]},
		mgl32.Vec3{center[0], center[1], center[2]},
		mgl32.Vec3{up[0], up[1], up[2]},
	)

	lookAtUniform := gl.GetUniformLocation(
		obj.prog.Identifier(),
		gl.Str(lookAtVariable),
	)

	gl.UniformMatrix4fv(
		lookAtUniform,
		1,
		false,
		&lookAtMat[0],
	)

	// rotate then translate:
	obj.angle += float32(delta.Seconds()) // * float32(orientation[3]*math.Pi/180)
	fmt.Printf("\n%f\n", obj.angle)
	rorateMat := mgl32.HomogRotate3D(obj.angle, mgl32.Vec3{orientation[0], orientation[1], orientation[2]})

	// translate:
	//transMat := mgl32.Translate3D(pos[0], pos[1], pos[2])

	// model matrix:
	//modelMat := mgl32.Ident4() //transMat.Add(rorateMat)

	// apply the model matrix:
	uniform := obj.UniformVariable()
	gl.UniformMatrix4fv(uniform, 1, false, &rorateMat[0])

	// vao:
	vao := obj.VAO()
	gl.BindVertexArray(vao)

	// render the material:
	texture, err := renderApp.Render(obj.material)
	if err != nil {
		return err
	}

	// render the material:
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.Identifier())

	// draw:
	amount := obj.VertexAmount()
	gl.DrawArrays(gl.TRIANGLES, 0, amount)

	return nil
}
