package opengl

import (
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

// NewApplication creates a new gui OpenGL application
func NewApplication(
	windowBuilder application_window.Builder,
) gui.Application {
	programsApp := programs.NewApplication()
	return createApplication(
		windowBuilder,
		programsApp,
	)
}
