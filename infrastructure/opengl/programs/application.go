package programs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	domain_materials "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	domain_layers "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	domain_layer "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/materials/material"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/shaders"
)

type application struct {
	programBuilder     program.Builder
	programsBuilder    Builder
	materialBuilder    material.Builder
	materialsBuilder   materials.Builder
	layerBuilder       layer.Builder
	layersBuilder      layers.Builder
	shadersApplication shaders.Application
}

func createApplication(
	programBuilder program.Builder,
	programsBuilder Builder,
	materialBuilder material.Builder,
	materialsBuilder materials.Builder,
	layerBuilder layer.Builder,
	layersBuilder layers.Builder,
	shadersApplication shaders.Application,
) Application {
	out := application{
		programBuilder:     programBuilder,
		programsBuilder:    programsBuilder,
		materialBuilder:    materialBuilder,
		materialsBuilder:   materialsBuilder,
		layerBuilder:       layerBuilder,
		layersBuilder:      layersBuilder,
		shadersApplication: shadersApplication,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(world worlds.World) (Programs, error) {
	programs := []program.Program{}
	scenes := world.Scenes()
	for _, oneScene := range scenes {
		prog, err := app.scene(oneScene)
		if err != nil {
			return nil, err
		}

		if prog == nil {
			continue
		}

		programs = append(programs, prog)
	}

	return app.programsBuilder.Create().WithPrograms(programs).Now()
}

func (app *application) scene(scene scenes.Scene) (program.Program, error) {
	if !scene.HasNodes() {
		fmt.Printf("\n**%s\n", scene.Hash().String())
		return nil, nil
	}

	nodes := scene.Nodes()
	compiledMaterialsList, err := app.node(nodes)
	if err != nil {
		return nil, err
	}

	compiledMaterials, err := app.materialsBuilder.Create().WithMaterials(compiledMaterialsList).Now()
	if err != nil {
		return nil, err
	}

	// create program:
	program := gl.CreateProgram()

	// attach all compiled shaders:
	compiledShaders := compiledMaterials.CompiledShaders()
	for _, oneCompiledShader := range compiledShaders {
		gl.AttachShader(program, oneCompiledShader.Identifier())
	}

	// link the program:
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		str := fmt.Sprintf("failed to link program: %s", log)
		return nil, errors.New(str)
	}

	// delete compiled shaders:
	for _, oneCompiledShader := range compiledShaders {
		gl.DeleteShader(oneCompiledShader.Identifier())
	}

	sceneHash := scene.Hash()
	return app.programBuilder.Create().
		WithScene(sceneHash).
		WithCompiledMaterials(compiledMaterials).
		WithIdentifier(program).
		Now()
}

func (app *application) node(nodes []nodes.Node) ([]material.Material, error) {
	materials := []material.Material{}
	for _, oneNode := range nodes {
		if oneNode.HasContent() {
			content := oneNode.Content()
			material, err := app.content(content)
			if err != nil {
				return nil, err
			}

			if material == nil {
				continue
			}

			materials = append(materials, material)
		}

		if oneNode.HasChildren() {
			children := oneNode.Children()
			childrenMaterials, err := app.node(children)
			if err != nil {
				return nil, err
			}

			for _, oneChildrenMaterial := range childrenMaterials {
				materials = append(materials, oneChildrenMaterial)
			}
		}
	}

	return materials, nil
}

func (app *application) content(content nodes.Content) (material.Material, error) {
	if content.IsModel() {
		material := content.Model().Material()
		return app.material(material)
	}

	return nil, nil
}

func (app *application) material(material domain_materials.Material) (material.Material, error) {
	layers := material.Layers()
	compiledLayers, err := app.layers(layers)
	if err != nil {
		return nil, err
	}

	hash := material.Hash()
	return app.materialBuilder.
		Create().
		WithCompiledLayers(compiledLayers).
		WithMaterial(hash).
		Now()
}

func (app *application) layers(layers domain_layers.Layers) (layers.Layers, error) {
	all := layers.All()
	compiledLayers := []layer.Layer{}
	for _, oneLayer := range all {
		compiledLayer, err := app.layer(oneLayer)
		if err != nil {
			return nil, err
		}

		compiledLayers = append(compiledLayers, compiledLayer)
	}

	return app.layersBuilder.Create().WithCompiledLayers(compiledLayers).Now()
}

func (app *application) layer(layer domain_layer.Layer) (layer.Layer, error) {
	shaders := layer.Shaders()
	compiledShaders, err := app.shadersApplication.Compile(shaders)
	if err != nil {
		return nil, err
	}

	hash := layer.Hash()
	return app.layerBuilder.Create().
		WithLayer(hash).
		WithCompiledShaders(compiledShaders).
		Now()
}
