package miners

import (
	"time"

	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/domain/memory/blocks"
	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type blockMiner struct {
	hashAdapter       hash.Adapter
	blockBuilder      blocks.Builder
	minedBlockBuilder mined_blocks.Builder
	identityService   identities.Service
	application       app.Application
	name              string
	seed              string
	password          string
	waitPeriod        time.Duration
	isStarted         bool
}

func createBlockMiner(
	hashAdapter hash.Adapter,
	blockBuilder blocks.Builder,
	minedBlockBuilder mined_blocks.Builder,
	identityService identities.Service,
	application app.Application,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
	isStarted bool,
) Application {
	out := blockMiner{
		hashAdapter:       hashAdapter,
		blockBuilder:      blockBuilder,
		minedBlockBuilder: minedBlockBuilder,
		identityService:   identityService,
		application:       application,
		name:              name,
		seed:              seed,
		password:          password,
		waitPeriod:        waitPeriod,
		isStarted:         isStarted,
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
		queueBuckets := identity.Wallet().Queue().All()

		// make the list of block buckets:
		blockBuckets := []buckets.Bucket{}
		for _, oneQueuedBucket := range queueBuckets {
			// add the bucket to the block list:
			blockBuckets = append(blockBuckets, oneQueuedBucket.Bucket())

			// remove the queued bucket from the queued identity:
			identity.Wallet().Broadcasted().Add(oneQueuedBucket)
		}

		// retrieve the chain:
		chain, err := app.application.Sub().Chain().Retrieve()
		if err != nil {
			return err
		}

		// calculate the difficulty:
		difficulty := difficulty(chain, uint(len(queueBuckets)))

		// build the block:
		createdOn := time.Now().UTC()
		gen := chain.Genesis()
		block, err := app.blockBuilder.Create().
			WithGenesis(gen).
			WithBuckets(blockBuckets).
			CreatedOn(createdOn).
			Now()

		if err != nil {
			return err
		}

		// mine the block:
		minedCreatedOn := time.Now().UTC()
		results, err := mine(app.hashAdapter, difficulty, block.Hash())
		if err != nil {
			return err
		}

		minedBlock, err := app.minedBlockBuilder.Create().
			WithBlock(block).
			WithMining(results).
			CreatedOn(minedCreatedOn).
			Now()

		if err != nil {
			return err
		}

		// add the block to the wallet:
		err = identity.Wallet().Blocks().Add(minedBlock)
		if err != nil {
			return err
		}

		// save the identity:
		err = app.identityService.Update(identity.Hash(), identity, app.password, app.password)
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
