package cameras

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs/program"
)

// NewApplication creates a new camera application
func NewApplication(
	currentSceneIndex uint,
	currentCameraIndex uint,
) Application {
	return createApplication(currentSceneIndex, currentCameraIndex)
}

// Application represents the camera application
type Application interface {
	Init(program program.Program, world worlds.World) error
}
