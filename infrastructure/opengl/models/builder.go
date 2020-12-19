package models

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type builder struct {
	materialBuilder materials.Builder
	model           models.Model
	prog            programs.Program
}

func createBuilder(
	materialBuilder materials.Builder,
) Builder {
	out := builder{
		materialBuilder: materialBuilder,
		model:           nil,
		prog:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.materialBuilder)
}

// WithModel adds a model to the builder
func (app *builder) WithModel(model models.Model) Builder {
	app.model = model
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Model instance
func (app *builder) Now() (Model, error) {
	if app.model == nil {
		return nil, errors.New("the model is mandatory in order to build a Model instance")
	}

	if app.prog == nil {
		return nil, errors.New("the program is mandatory in order to build a Model instance")
	}

	domainMaterial := app.model.Material()
	material, err := app.materialBuilder.Create().WithMaterial(domainMaterial).WithProgram(app.prog).Now()
	if err != nil {
		return nil, err
	}

	domainGeo := app.model.Geometry()
	vao, vertexAmount, isTriangle, err := app.geometry(app.prog, domainGeo)
	if err != nil {
		return nil, err
	}

	modelMat := mgl32.Ident4()
	identifier := app.prog.Identifier()
	varName := fmt.Sprintf(glStrPattern, app.model.Variable())
	modelUniform := gl.GetUniformLocation(identifier, gl.Str(varName))
	gl.UniformMatrix4fv(modelUniform, 1, false, &modelMat[0])

	// create the type:
	var typ Type
	if isTriangle {
		typ = createTypeWithTriangle()
	}

	if typ == nil {
		return nil, errors.New("the type (isTriangle) is mandatory in order to build a Model instance")
	}

	return createModel(app.model, typ, vao, int32(vertexAmount), modelUniform, app.prog, material), nil
}

func (app *builder) geometry(program programs.Program, geometry geometries.Geometry) (uint32, int, bool, error) {
	vertices := geometry.Vertices()
	variables := geometry.Variables()
	vertexCoordinatesVariable := variables.VertexCoordinates()
	textureCoordinatesVariable := variables.TextureCoordinates()

	list := []float32{}
	all := vertices.All()
	for _, oneVertice := range all {
		pos := oneVertice.Position()
		tex := oneVertice.Texture()
		list = append(list, []float32{
			pos.X(),
			pos.Y(),
			pos.Z(),
			tex.X(),
			tex.Y(),
		}...)
	}

	verticesType := vertices.Type()
	if verticesType.IsTriangle() {
		triSize := int32(3)
		texSize := int32(2)
		vertexSize := triSize + texSize
		stride := int32(vertexSize * float32SizeInBytes)
		identifier := program.Identifier()

		var vao uint32
		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)

		var vbo uint32
		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(list)*float32SizeInBytes, gl.Ptr(list), gl.STATIC_DRAW)

		verOffset := int(0)
		vertexVarName := fmt.Sprintf(glStrPattern, vertexCoordinatesVariable)
		vertAttrib := uint32(gl.GetAttribLocation(identifier, gl.Str(vertexVarName)))
		gl.EnableVertexAttribArray(vertAttrib)
		gl.VertexAttribPointer(vertAttrib, triSize, gl.FLOAT, false, stride, gl.PtrOffset(verOffset))

		texOffset := int(triSize * float32SizeInBytes)
		texVarName := fmt.Sprintf(glStrPattern, textureCoordinatesVariable)
		texCoordAttrib := uint32(gl.GetAttribLocation(identifier, gl.Str(texVarName)))
		gl.EnableVertexAttribArray(texCoordAttrib)
		gl.VertexAttribPointer(texCoordAttrib, texSize, gl.FLOAT, false, stride, gl.PtrOffset(texOffset))

		// return:
		return vao, len(list), true, nil
	}

	return 0, 0, false, errors.New("the vertices type is invalid")
}
