package opengl

import (
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/worlds"
)

// NewApplication creates a new gui OpenGL application
func NewApplication(
	windowBuilder application_window.Builder,
	defaultSceneIndex uint,
	defaultCameraIndex uint,
) gui.Application {
	worldBuilder := worlds.NewBuilder(
		defaultSceneIndex,
		defaultCameraIndex,
	)

	return createApplication(
		windowBuilder,
		worldBuilder,
		defaultSceneIndex,
	)
}
