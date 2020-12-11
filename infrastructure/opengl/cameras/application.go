package cameras

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
)

type application struct {
	currentSceneIndex  uint
	currentCameraIndex uint
}

func createApplication(
	currentSceneIndex uint,
	currentCameraIndex uint,
) Application {
	out := application{
		currentSceneIndex:  currentSceneIndex,
		currentCameraIndex: currentCameraIndex,
	}

	return &out
}

// Init initializes the application
func (app *application) Init(program program.Program, world worlds.World) error {
	// fetch the current scene:
	scene, err := world.Scene(app.currentSceneIndex)
	if err != nil {
		return err
	}

	// fetch the current camera:
	cam, err := scene.Camera(app.currentCameraIndex)
	if err != nil {
		return err
	}

	// fetch the program identifier:
	progIdentifier := program.Identifier()

	// projection:
	projection := cam.Projection()
	projVariable := fmt.Sprintf("%s\x00", projection.Variable())
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

	// camera:
	lookAt := cam.LookAt()
	lookAtVariable := fmt.Sprintf("%s\x00", lookAt.Variable())
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

	return nil
}
