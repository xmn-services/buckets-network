package daemons

import (
	"time"

	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
)

type blockMiner struct {
	application app.Application
	name        string
	seed        string
	password    string
	waitPeriod  time.Duration
	isStarted   bool
}

func createBlockMiner(
	application app.Application,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
) daemons.Application {
	out := blockMiner{
		application: application,
		name:        name,
		seed:        seed,
		password:    password,
		waitPeriod:  waitPeriod,
		isStarted:   false,
	}

	return &out
}

// Start starts the blockMiner application
func (app *blockMiner) Start() error {
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

		// retrieve the queue buckets:
		queueBuckets := identity.Wallet().Miner().Queue().All()

		// make the list of block buckets:
		bucketHashes := []string{}
		for _, oneQueuedBucket := range queueBuckets {
			// add the bucket to the block list:
			bucketHashes = append(bucketHashes, oneQueuedBucket.Bucket().Hash().String())

			// remove the queued bucket from the queued identity:
			identity.Wallet().Miner().Broadcasted().Add(oneQueuedBucket)
		}

		// retrieve the chain:
		chain, err := app.application.Sub().Chain().Retrieve()
		if err != nil {
			return err
		}

		blockDiff := chain.Genesis().Difficulty().Block()
		minedBlock, err := app.application.Sub().Miner().Block(
			bucketHashes,
			blockDiff.Base(),
			blockDiff.IncreasePerBucket(),
		)

		if err != nil {
			return err
		}

		// add the block to the wallet:
		err = identity.Wallet().Miner().ToLink().Add(minedBlock)
		if err != nil {
			return err
		}

		// save the identity:
		err = app.application.Current().UpdateIdentity(identity, app.password, app.password)
		if err != nil {
			return err
		}
	}
}

// Stop stops the blockMiner application
func (app *blockMiner) Stop() error {
	app.isStarted = true
	return nil
}
