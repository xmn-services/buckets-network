package opengl

import (
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

// NewApplication creates a new gui OpenGL application
func NewApplication(
	windowBuilder application_window.Builder,
	currentSceneIndex uint,
	currentCameraIndex uint,
) gui.Application {
	programsApp := programs.NewApplication()
	cameraApp := cameras.NewApplication(currentSceneIndex, currentCameraIndex)
	return createApplication(
		windowBuilder,
		programsApp,
		cameraApp,
		currentSceneIndex,
	)
}
