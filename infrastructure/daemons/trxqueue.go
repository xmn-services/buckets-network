package daemons

import (
	"time"

	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
)

type trxQueue struct {
	application app.Application
	name        string
	seed        string
	password    string
	waitPeriod  time.Duration
	isStarted   bool
}

func createQueue(
	app app.Application,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
	isStarted bool,
) daemons.Application {
	out := trxQueue{
		application: app,
		name:        name,
		seed:        seed,
		password:    password,
		waitPeriod:  waitPeriod,
		isStarted:   isStarted,
	}

	return &out
}

// Start starts the application
func (app *trxQueue) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the identity:
		identityApp, err := app.application.Current().Authenticate(app.name, app.seed, app.password)
		if err != nil {
			return err
		}

		identity, err := identityApp.Current().Retrieve()
		if err != nil {
			return err
		}

		// retrieve the buckets that have not been added to a transaction yet:
		buckets := identity.Wallet().Miner().ToTransact().All()

		// for each bucket, create a transaction:
		for _, oneBucket := range buckets {
			// add the transaction to the to-mine trxQueue:
			err = identity.Wallet().Miner().Queue().Add(oneBucket)
			if err != nil {
				return err
			}

			// remove the bucket from the new list:
			err := identity.Wallet().Miner().ToTransact().Delete(oneBucket.Hash())
			if err != nil {
				return err
			}
		}

		// update the identity:
		err = app.application.Current().UpdateIdentity(identity, app.password, app.password)
		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *trxQueue) Stop() error {
	app.isStarted = true
	return nil
}