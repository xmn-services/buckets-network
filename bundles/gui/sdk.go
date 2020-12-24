package gui

import (
	"github.com/xmn-services/buckets-network/application/gui"
	"github.com/xmn-services/buckets-network/application/windows"
	domain_colors "github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	domain_textures "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	domain_pixels "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/pixels"
	"github.com/xmn-services/buckets-network/infrastructure/glfw"
	"github.com/xmn-services/buckets-network/infrastructure/opengl"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders/applications"
)

// NewOpenglApplication creates a new OpenGL application
func NewOpenglApplication(
	currentSceneIndex uint,
	currentCameraIndex uint,
) gui.Application {
	colorBuilder := domain_colors.NewBuilder()
	pixelBuilder := domain_pixels.NewBuilder()
	textureBuilder := domain_textures.NewBuilder()
	renderAppBuilder := applications.NewBuilder(
		colorBuilder,
		pixelBuilder,
		textureBuilder,
	)

	builder := NewGlfwApplicationBuilder()
	return opengl.NewApplication(
		renderAppBuilder,
		builder,
		currentSceneIndex,
		currentCameraIndex,
	)
}

// NewGlfwApplicationBuilder creates a new glfw application builder
func NewGlfwApplicationBuilder() windows.Builder {
	return glfw.NewBuilder()
}
