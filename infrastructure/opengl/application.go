package opengl

import (
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/xmn-services/buckets-network/application/gui"
	application_window "github.com/xmn-services/buckets-network/application/windows"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	domain_worlds "github.com/xmn-services/buckets-network/domain/memory/worlds"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/worlds"
)

type application struct {
	windowBuilder application_window.Builder
	worldBuilder  worlds.Builder
	sceneIndex    uint
	currentWorld  worlds.World
}

func createApplication(
	windowBuilder application_window.Builder,
	worldBuilder worlds.Builder,
	sceneIndex uint,
) gui.Application {
	out := application{
		windowBuilder: windowBuilder,
		worldBuilder:  worldBuilder,
		sceneIndex:    sceneIndex,
		currentWorld:  nil,
	}

	return &out
}

// Execute executes a gui OpenGL application
func (app *application) Execute(win windows.Window, domainWorld domain_worlds.World) error {
	// create the window:
	winApp, err := app.windowBuilder.Create().WithWindow(win).Now()
	if err != nil {
		return err
	}

	// build the world:
	world, err := app.worldBuilder.Create().WithWorld(domainWorld).Now()
	if err != nil {
		return err
	}

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.5, 0.2, 0.4, 1.0)

	// set the world:
	app.currentWorld = world

	// execute the window app:
	return winApp.Execute(app.updateFn)
}

func (app *application) updateFn(prev time.Time, current time.Time) error {
	// render the world:
	elapsed := current.Sub(prev)
	err := app.currentWorld.Render(elapsed)
	if err != nil {
		return err
	}

	//log.Printf("\nupdate: %s -- %s", prev.String(), current.String())
	time.Sleep(time.Second / 20)
	return nil
}
