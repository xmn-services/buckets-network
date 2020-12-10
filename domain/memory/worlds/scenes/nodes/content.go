package nodes

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type content struct {
	model  models.Model
	camera cameras.Camera
}

func createContentWithModel(
	model models.Model,
) Content {
	return createContentInternally(model, nil)
}

func createContentWithCamera(
	camera cameras.Camera,
) Content {
	return createContentInternally(nil, camera)
}

func createContentInternally(
	model models.Model,
	camera cameras.Camera,
) Content {
	out := content{
		model:  model,
		camera: camera,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	if obj.IsModel() {
		return obj.Model().Hash()
	}

	return obj.Camera().Hash()
}

// IsModel returns true if the content is a model, false otherwise
func (obj *content) IsModel() bool {
	return obj.model != nil
}

// Model returns the model, if any
func (obj *content) Model() models.Model {
	return obj.model
}

// IsCamera returns true if the content is a camera, false otherwise
func (obj *content) IsCamera() bool {
	return obj.camera != nil
}

// Camera returns the camera, if any
func (obj *content) Camera() cameras.Camera {
	return obj.camera
}
