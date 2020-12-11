package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/xmn-services/buckets-network/bundles/gui"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels/pixel"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
)

func main() {
	// create the window:
	title := "My Window"
	width := uint(800)
	height := uint(600)
	window, err := windows.NewBuilder().Create().WithTitle(title).WithWidth(width).WithHeight(height).Now()
	if err != nil {
		panic(err)
	}

	// create the camera:
	fov := float64(0.5)
	index := uint(0)
	cameraNode := nodeFromCamera(camera(fov, index))

	// create the cube model:
	cubeModel := nodeFromModel(cubeModel())

	// create the scene:
	scene, err := scenes.NewFactory().Create()
	if err != nil {
		panic(err)
	}

	// add the camera:
	err = scene.Add(cameraNode)
	if err != nil {
		panic(err)
	}

	// add the cube:
	err = scene.Add(cubeModel)
	if err != nil {
		panic(err)
	}

	// create the world:
	world, err := worlds.NewFactory().Create()
	if err != nil {
		panic(err)
	}

	// add the scene to the world:
	err = world.Add(scene)
	if err != nil {
		panic(err)
	}

	app := gui.NewOpenglApplication()
	err = app.Execute(window, world)
	if err != nil {
		log.Println(err.Error())
	}
}

func nodeFromCamera(camera cameras.Camera) nodes.Node {
	pos := math.Vec3{0.0, 0.0, 0.0}
	right := math.Vec3{0.0, 0.0, 0.0}
	up := math.Vec3{0.0, 1.0, 0.0}
	node, err := nodes.NewBuilder().Create().WithPosition(pos).WithRight(right).WithUp(up).WithCamera(camera).Now()
	if err != nil {
		panic(err)
	}

	return node
}

func nodeFromModel(model models.Model) nodes.Node {
	pos := math.Vec3{1.0, 0.0, 0.0}
	right := math.Vec3{0.0, 0.0, 0.0}
	up := math.Vec3{0.0, 1.0, 0.0}
	node, err := nodes.NewBuilder().Create().WithPosition(pos).WithRight(right).WithUp(up).WithModel(model).Now()
	if err != nil {
		panic(err)
	}

	return node
}

func camera(fov float64, index uint) cameras.Camera {
	pos := math.Vec2{0.0, 0.0}
	dimension := math.Vec2{1.0, 1.0}

	rectangleBuilder := rectangles.NewBuilder()
	viewport, err := rectangleBuilder.Create().WithPosition(pos).WithDimension(dimension).Now()
	if err != nil {
		panic(err)
	}

	createdOn := time.Now().UTC()
	cameraBuilder := cameras.NewBuilder()
	camera, err := cameraBuilder.Create().WithViewport(viewport).WithFieldOfView(fov).WithIndex(index).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return camera
}

func cubeModel() models.Model {
	geo := cubeGeometry()
	mat := cubeMaterial()
	model, err := models.NewBuilder().Create().WithGeometry(geo).WithMaterial(mat).Now()
	if err != nil {
		panic(err)
	}

	return model
}

func cubeGeometry() geometries.Geometry {
	list := []vertex.Vertex{}
	amount := len(cubeVertices) / 5
	vertexBuilder := vertex.NewBuilder()
	for i := 0; i < amount; i++ {
		pos := math.Vec3{
			cubeVertices[i],
			cubeVertices[i+1],
			cubeVertices[i+2],
		}

		tex := math.Vec2{
			cubeVertices[i+3],
			cubeVertices[i+4],
		}

		vertex, err := vertexBuilder.Create().WithPosition(pos).WithTexture(tex).Now()
		if err != nil {
			panic(err)
		}

		list = append(list, vertex)
	}

	vertices, err := vertices.NewBuilder().Create().WithoutHash().WithVertices(list).Now()
	if err != nil {
		panic(err)
	}

	shaders := cubeVertexShader()
	geo, err := geometries.NewBuilder().Create().WithShaders(shaders).WithVertices(vertices).Now()
	if err != nil {
		panic(err)
	}

	return geo
}

func cubeMaterial() materials.Material {
	pos := math.Vec2{0.0, 0.0}
	dimension := math.Vec2{1.0, 1.0}
	viewport, err := rectangles.NewBuilder().Create().WithPosition(pos).WithDimension(dimension).Now()
	if err != nil {
		panic(err)
	}

	tex := generateTexture()
	rdn, err := render.NewBuilder().Create().WithOpacity(1.0).WithViewport(viewport).WithTexture(tex).Now()
	if err != nil {
		panic(err)
	}

	rdns, err := renders.NewBuilder().Create().WithoutHash().WithRenders([]render.Render{
		rdn,
	}).Now()

	if err != nil {
		panic(err)
	}

	shades := cubeFragmentShader()
	lyr, err := layer.NewBuilder().Create().WithOpacity(1.0).WithViewport(viewport).WithRenders(rdns).WithShaders(shades).Now()
	if err != nil {
		panic(err)
	}

	layerList, err := layers.NewBuilder().Create().WithoutHash().WithLayers([]layer.Layer{
		lyr,
	}).Now()

	if err != nil {
		panic(err)
	}

	mat, err := materials.NewBuilder().Create().WithOpacity(1.0).WithViewport(viewport).WithLayers(layerList).Now()
	if err != nil {
		panic(err)
	}

	return mat
}

func cubeVertexShader() shaders.Shaders {
	shaderBuilder := shader.NewBuilder()
	vShader, err := shaderBuilder.WithCode(vertexShader).WithVariables([]string{
		"projection",
		"camera",
		"model",
	}).IsVertex().Now()

	if err != nil {
		panic(err)
	}

	out, err := shaders.NewBuilder().WithoutHash().WithShaders([]shader.Shader{
		vShader,
	}).Now()
	if err != nil {
		panic(err)
	}

	return out
}

func cubeFragmentShader() shaders.Shaders {
	shaderBuilder := shader.NewBuilder()
	fShader, err := shaderBuilder.WithCode(fragmentShader).WithVariables([]string{
		"tex",
	}).IsFragment().Now()

	if err != nil {
		panic(err)
	}

	out, err := shaders.NewBuilder().WithoutHash().WithShaders([]shader.Shader{
		fShader,
	}).Now()
	if err != nil {
		panic(err)
	}

	return out
}

func generateTexture() textures.Texture {
	pos := math.Vec2{0.0, 0.0}
	dim := math.Vec2{500.0, 500.0}
	dimension, err := rectangles.NewBuilder().Create().WithPosition(pos).WithDimension(dim).Now()
	if err != nil {
		panic(err)
	}

	colorBuilder := colors.NewBuilder()
	pixelBuilder := pixel.NewBuilder()
	pixelsBuilder := pixels.NewBuilder()

	width := 500
	height := 500
	alpha := uint32(1)
	rowsList := []pixels.Pixels{}
	for i := 0; i < width; i++ {
		list := []pixel.Pixel{}
		for j := 0; j < height; j++ {
			red := rand.Uint32()
			green := rand.Uint32()
			blue := rand.Uint32()
			color := colorBuilder.Create().WithRed(red).WithGreen(green).WithBlue(blue).Now()
			pixel, err := pixelBuilder.Create().WithColor(color).WithAlpha(alpha).Now()
			if err != nil {
				panic(err)
			}

			list = append(list, pixel)
		}

		pixels, err := pixelsBuilder.Create().WithoutHash().WithPixels(list).Now()
		if err != nil {
			panic(err)
		}

		rowsList = append(rowsList, pixels)
	}

	rws, err := rows.NewBuilder().WithoutHash().WithRows(rowsList).Now()
	if err != nil {
		panic(err)
	}

	tex, err := textures.NewBuilder().Create().WithDimension(dimension).WithPixels(rws).Now()
	if err != nil {
		panic(err)
	}

	return tex
}

var vertexShader = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

var cubeVertices = []float64{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}
