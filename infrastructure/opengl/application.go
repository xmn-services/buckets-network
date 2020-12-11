package opengl

import (
	"errors"
	"log"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type application struct {
	windowBuilder application_window.Builder
	programsApp   programs.Application
}

func createApplication(
	windowBuilder application_window.Builder,
	programsApp programs.Application,
) gui.Application {
	out := application{
		windowBuilder: windowBuilder,
		programsApp:   programsApp,
	}

	return &out
}

// Execute executes a gui OpenGL application
func (app *application) Execute(win windows.Window, world worlds.World) error {
	if world == nil {
		return errors.New("the world is mandatory in order to execute the application")
	}

	winApp, err := app.windowBuilder.Create().WithWindow(win).Now()
	if err != nil {
		return err
	}

	// init:
	err = app.init(world)
	if err != nil {
		return err
	}

	return winApp.Execute(app.updateFn)
}

func (app *application) init(world worlds.World) error {
	// initialize OpenGL:
	err := gl.Init()
	if err != nil {
		return err
	}

	// log the OpenGL version:
	version := gl.GoStr(gl.GetString(gl.VERSION))
	if version == "" {
		return errors.New("there was an error during the initialization of OpenGL, since the version could not be fetched")
	}

	// log the OpenGL version:
	log.Printf("\nOpenGL version: %s", version)

	// compile the program:
	program, err := app.programsApp.Execute(world)
	if err != nil {
		return err
	}

	log.Printf("\n****program: %v\n", program)

	// returns:
	return nil
}

func (app *application) updateFn(prev time.Time, current time.Time) error {
	//log.Printf("\nupdate: %s -- %s", prev.String(), current.String())
	time.Sleep(time.Second / 20)
	return nil
}
