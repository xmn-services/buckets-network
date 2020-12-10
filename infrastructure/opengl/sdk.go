package opengl

import (
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
)

// NewApplication creates a new gui OpenGL application
func NewApplication(
	windowBuilder application_window.Builder,
) gui.Application {
	return createApplication(windowBuilder)
}
