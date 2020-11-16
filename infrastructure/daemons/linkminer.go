package daemons

import (
	"time"

	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
)

type linkMiner struct {
	application app.Application
	name        string
	seed        string
	password    string
	waitPeriod  time.Duration
	isStarted   bool
}

func createLinkMiner(
	application app.Application,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
) daemons.Application {
	out := linkMiner{
		application: application,
		name:        name,
		seed:        seed,
		password:    password,
		waitPeriod:  waitPeriod,
		isStarted:   false,
	}

	return &out
}

// Start starts the linkMiner application
func (app *linkMiner) Start() error {
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

		// retrieve the chain:
		chain, err := app.application.Sub().Chain().Retrieve()
		if err != nil {
			return err
		}

		// fetch the mined blocks to link:
		blocks := identity.Wallet().Miner().ToLink().All()
		for _, oneBlock := range blocks {
			// mine the link:
			difficulty := chain.Genesis().Difficulty().Link()
			minedLink, err := app.application.Sub().Miner().Link(
				chain.Head().Hash().String(),
				oneBlock.Hash().String(),
				difficulty,
			)

			if err != nil {
				return err
			}

			// update the chain:
			err = app.application.Sub().Chain().Update(minedLink)
			if err != nil {
				return err
			}
		}

	}
}

// Stop stops the linkMiner application
func (app *linkMiner) Stop() error {
	app.isStarted = true
	return nil
}
