package nodes

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/models"
)

type content struct {
	model models.Model
	cam   cameras.Camera
}

func createContentWithModel(
	model models.Model,
) Content {
	return createContentInternally(model, nil)
}

func createContentWithCamera(
	cam cameras.Camera,
) Content {
	return createContentInternally(nil, cam)
}

func createContentInternally(
	model models.Model,
	cam cameras.Camera,
) Content {
	out := content{
		model: model,
		cam:   cam,
	}

	return &out
}

// IsModel returns true if there is a model, false otherwise
func (obj *content) IsModel() bool {
	return obj.model != nil
}

// Model returns the model, if any
func (obj *content) Model() models.Model {
	return obj.model
}

// IsCamera returns true if there is a camera, false otherwise
func (obj *content) IsCamera() bool {
	return obj.cam != nil
}

// Camera returns the camera, if any
func (obj *content) Camera() cameras.Camera {
	return obj.cam
}
