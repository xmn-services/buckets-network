package gui

import (
	"github.com/xmn-services/buckets-network/application/gui"
	"github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/infrastructure/glfw"
	"github.com/xmn-services/buckets-network/infrastructure/opengl"
)

// NewOpenglApplication creates a new OpenGL application
func NewOpenglApplication(
	currentSceneIndex uint,
	currentCameraIndex uint,
) gui.Application {
	builder := NewGlfwApplicationBuilder()
	return opengl.NewApplication(
		builder,
		currentSceneIndex,
		currentCameraIndex,
	)
}

// NewGlfwApplicationBuilder creates a new glfw application builder
func NewGlfwApplicationBuilder() windows.Builder {
	return glfw.NewBuilder()
}
