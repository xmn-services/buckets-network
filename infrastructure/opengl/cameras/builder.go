package cameras

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type builder struct {
	program programs.Program
	camera  cameras.Camera
}

func createBuilder() Builder {
	out := builder{
		program: nil,
		camera:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.program = prog
	return app
}

// WithCamera adds a camera to the builder
func (app *builder) WithCamera(camera cameras.Camera) Builder {
	app.camera = camera
	return app
}

// Now builds a new Camera instance
func (app *builder) Now() (Camera, error) {
	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Camera instance")
	}

	if app.camera == nil {
		return nil, errors.New("the camera is mandatory in order ro build a Camera instance")
	}

	// fetch the program identifier:
	progIdentifier := app.program.Identifier()

	// projection:
	projection := app.camera.Projection()
	projVariable := fmt.Sprintf(glStrVarPattern, projection.Variable())
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
		progIdentifier,
		gl.Str(projVariable),
	)

	gl.UniformMatrix4fv(
		projectionUniform,
		1,
		false,
		&projMat[0],
	)

	proj := createMatrix(projectionUniform, projMat)

	// camera:
	lookAt := app.camera.LookAt()
	lookAtVariable := fmt.Sprintf(glStrVarPattern, lookAt.Variable())
	eye := lookAt.Eye()
	center := lookAt.Center()
	up := lookAt.Up()

	camera := mgl32.LookAtV(
		mgl32.Vec3{eye[0], eye[1], eye[2]},
		mgl32.Vec3{center[0], center[1], center[2]},
		mgl32.Vec3{up[0], up[1], up[2]},
	)

	cameraUniform := gl.GetUniformLocation(
		progIdentifier,
		gl.Str(lookAtVariable),
	)

	gl.UniformMatrix4fv(
		cameraUniform,
		1,
		false,
		&camera[0],
	)

	position := createMatrix(cameraUniform, camera)
	return createCamera(app.camera, proj, position), nil
}
