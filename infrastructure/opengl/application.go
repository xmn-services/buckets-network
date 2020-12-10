package opengl

import (
	"log"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
)

type application struct {
	windowBuilder application_window.Builder
}

func createApplication(
	windowBuilder application_window.Builder,
) gui.Application {
	out := application{
		windowBuilder: windowBuilder,
	}

	return &out
}

// Execute executes a gui OpenGL application
func (app *application) Execute(win windows.Window, world worlds.World) error {
	winApp, err := app.windowBuilder.Create().WithWindow(win).Now()
	if err != nil {
		return err
	}

	// init:
	err = app.init()
	if err != nil {
		return err
	}

	return winApp.Execute(app.updateFn)
}

func (app *application) init() error {
	// initialize OpenGL:
	err := gl.Init()
	if err != nil {
		return err
	}

	// log the OpenGL version:
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("\nOpenGL version: %s", version)

	// returns:
	return nil
}

func (app *application) updateFn(prev time.Time, current time.Time) error {
	log.Printf("\nupdate: %s -- %s", prev.String(), current.String())
	time.Sleep(time.Second / 20)
	return nil
}
