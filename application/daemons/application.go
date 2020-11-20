package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application/syncs"
)

type application struct {
	syncApp    syncs.Application
	waitPeriod time.Duration
	isStarted  bool
}

func createApplication(
	syncApp syncs.Application,
	waitPeriod time.Duration,
) Application {
	out := application{
		syncApp:    syncApp,
		waitPeriod: waitPeriod,
		isStarted:  false,
	}

	return &out
}

// Start starts the application
func (app *application) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// download new the peers
		err := app.syncApp.Peers()
		if err != nil {
			// log
		}

		// download new head link if needed
		err = app.syncApp.Chain()
		if err != nil {
			// log
		}

		// mine the block and links:
		err = app.syncApp.Miners()
		if err != nil {
			// log
		}

		// download needed buckets:
		err = app.syncApp.Storage()
		if err != nil {
			// log
		}
	}
}

// Stop stops the application
func (app *application) Stop() error {
	app.isStarted = true
	return nil
}
