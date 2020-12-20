package main

import (
	"log"
	"time"

	"github.com/xmn-services/buckets-network/bundles/gui"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
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
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/shaders/shader"
)

func main() {
	// create the window:
	title := "My Window"
	width := uint(1600)
	height := uint(600)
	window, err := windows.NewBuilder().Create().WithTitle(title).WithWidth(width).WithHeight(height).Now()
	if err != nil {
		panic(err)
	}

	// create the camera:
	cameraIndex := uint(0)
	cameraNode := nodeFromCamera(camera(width, height, cameraIndex))

	// create the cube model:
	cubeModel := nodeFromModel(cubeModel())

	// create the world:
	world, err := worlds.NewFactory().Create()
	if err != nil {
		panic(err)
	}

	// fetch the scene:
	scene, err := world.Scene(scenes.CurrentSceneIndex)
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

	app := gui.NewOpenglApplication(
		scenes.CurrentSceneIndex,
		cameraIndex,
	)

	err = app.Execute(window, world)
	if err != nil {
		log.Println(err.Error())
	}
}

func nodeFromCamera(camera cameras.Camera) nodes.Node {
	pos := fl32.Vec3{0.0, 0.0, 0.0}
	angle := float32(45.0)
	direction := fl32.Vec3{5.0, 5.0, 5.0}
	node, err := nodes.NewBuilder().Create().WithPosition(pos).WithOrientationAngle(angle).WithOrientationDirection(direction).WithCamera(camera).Now()
	if err != nil {
		panic(err)
	}

	return node
}

func nodeFromModel(model models.Model) nodes.Node {
	pos := fl32.Vec3{1.0, 1.0, 1.0}
	angle := float32(200.0)
	direction := fl32.Vec3{0.0, 0.0, 1.0}
	node, err := nodes.NewBuilder().Create().WithPosition(pos).WithOrientationAngle(angle).WithOrientationDirection(direction).WithModel(model).Now()
	if err != nil {
		panic(err)
	}

	return node
}

func camera(width uint, height uint, index uint) cameras.Camera {
	lookAtVariable := "camera"
	eye := fl32.Vec3{5, 5, 5}
	center := fl32.Vec3{0, 0, 0}
	up := fl32.Vec3{0, 1, 0}

	projectionVariable := "projection"
	fov := float32(45.0)
	aspectRatio := float32(width / height)
	near := float32(0.1)
	far := float32(10.0)
	createdOn := time.Now().UTC()
	camera, err := cameras.NewBuilder().Create().
		WithLookAtVariable(lookAtVariable).
		WithLookAtEye(eye).
		WithLookAtCenter(center).
		WithLookAtUp(up).
		WithProjectionVariable(projectionVariable).
		WithProjectionFieldofView(fov).
		WithProjectionAspectRatio(aspectRatio).
		WithProjectionNear(near).
		WithProjectionFar(far).
		WithIndex(index).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return camera
}

func cubeModel() models.Model {
	geo := cubeGeometry()
	mat := cubeMaterial()
	variable := "model"
	model, err := models.NewBuilder().Create().WithGeometry(geo).WithMaterial(mat).WithVariable(variable).Now()
	if err != nil {
		panic(err)
	}

	return model
}

func cubeGeometry() geometries.Geometry {
	list := []vertex.Vertex{}
	amount := len(cubeVertices)
	vertexBuilder := vertex.NewBuilder()
	for i := 0; i < amount; i++ {
		pos := fl32.Vec3{
			cubeVertices[i],
			cubeVertices[i+1],
			cubeVertices[i+2],
		}

		tex := fl32.Vec2{
			cubeVertices[i+3],
			cubeVertices[i+4],
		}

		vertex, err := vertexBuilder.Create().WithPosition(pos).WithTexture(tex).Now()
		if err != nil {
			panic(err)
		}

		list = append(list, vertex)
		i += 4
	}

	vertices, err := vertices.NewBuilder().Create().WithoutHash().WithVertices(list).IsTriangle().Now()
	if err != nil {
		panic(err)
	}

	vertexCoordinatesVar := "vert"
	texCoordinatesVar := "vertTexCoord"
	shaders := cubeVertexShader()
	geo, err := geometries.NewBuilder().Create().
		WithVertices(vertices).
		WithShaders(shaders).
		WithVertexCoordinatesVariable(vertexCoordinatesVar).
		WithTextureCoordinatesVariable(texCoordinatesVar).
		Now()

	if err != nil {
		panic(err)
	}

	return geo
}

func cubeMaterial() materials.Material {
	pos := ints.Vec2{0.0, 0.0}
	dimension := ints.Vec2{1.0, 1.0}
	viewport, err := ints.NewBuilder().Create().WithPosition(pos).WithDimension(dimension).Now()
	if err != nil {
		panic(err)
	}

	tex := generateTexture()
	rdn, err := renders.NewBuilder().Create().WithOpacity(1.0).WithViewport(viewport).WithTexture(tex).Now()
	if err != nil {
		panic(err)
	}

	lyr, err := layer.NewBuilder().Create().WithAlpha(1.0).WithViewport(viewport).WithRender(rdn).Now()
	if err != nil {
		panic(err)
	}

	layerList, err := layers.NewBuilder().Create().WithoutHash().WithLayers([]layer.Layer{
		lyr,
	}).Now()

	if err != nil {
		panic(err)
	}

	shades := cubeFragmentShader()
	mat, err := materials.NewBuilder().Create().WithShaders(shades).WithAlpha(1.0).WithViewport(viewport).WithLayers(layerList).Now()
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
	pos := ints.Vec2{0, 0}
	dim := ints.Vec2{512, 512}
	viewport, err := ints.NewBuilder().Create().WithPosition(pos).WithDimension(dim).Now()
	if err != nil {
		panic(err)
	}

	colorBuilder := colors.NewBuilder()
	pixelBuilder := pixels.NewBuilder()

	width := pos.X() + dim.X()
	height := pos.X() + dim.Y()
	total := width * height
	alpha := uint8(255)
	pixels := []pixels.Pixel{}
	for i := 0; i < total; i++ {
		red := 0xff
		green := 0x00
		blue := 0x00
		color := colorBuilder.Create().WithRed(uint8(red)).WithGreen(uint8(green)).WithBlue(uint8(blue)).Now()
		pixel, err := pixelBuilder.Create().WithColor(color).WithAlpha(alpha).Now()
		if err != nil {
			panic(err)
		}

		pixels = append(pixels, pixel)
	}

	tex, err := textures.NewBuilder().Create().WithViewport(viewport).WithPixels(pixels).Now()
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

var cubeVertices = []float32{
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
