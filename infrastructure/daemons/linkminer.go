package daemons

import (
	"time"

	app "github.com/xmn-services/buckets-network/application"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/links"
	mined_links "github.com/xmn-services/buckets-network/domain/memory/links/mined"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type linkMiner struct {
	hashAdapter      hash.Adapter
	linkBuilder      links.Builder
	minedLinkBuilder mined_links.Builder
	chainBuilder     chains.Builder
	application      app.Application
	name             string
	seed             string
	password         string
	waitPeriod       time.Duration
	isStarted        bool
}

func createLinkMiner(
	hashAdapter hash.Adapter,
	linkBuilder links.Builder,
	minedLinkBuilder mined_links.Builder,
	chainBuilder chains.Builder,
	application app.Application,
	name string,
	seed string,
	password string,
	waitPeriod time.Duration,
	isStarted bool,
) daemons.Application {
	out := linkMiner{
		hashAdapter:      hashAdapter,
		linkBuilder:      linkBuilder,
		minedLinkBuilder: minedLinkBuilder,
		chainBuilder:     chainBuilder,
		application:      application,
		name:             name,
		seed:             seed,
		password:         password,
		waitPeriod:       waitPeriod,
		isStarted:        isStarted,
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
			prev := chain.Head().Hash()
			linkCreatedOn := time.Now().UTC()
			link, err := app.linkBuilder.Create().
				WithPreviousLink(prev).
				WithNext(oneBlock).
				CreatedOn(linkCreatedOn).
				Now()

			if err != nil {
				return err
			}

			// mine:
			difficulty := chain.Genesis().Difficulty().Link()
			results, err := mine(app.hashAdapter, difficulty, link.Hash())
			if err != nil {
				return err
			}

			// return the mined link:
			minedLinkCreatedOn := time.Now().UTC()
			minedLink, err := app.minedLinkBuilder.Create().
				WithLink(link).
				WithMining(results).
				CreatedOn(minedLinkCreatedOn).
				Now()

			if err != nil {
				return err
			}

			gen := chain.Genesis()
			root := chain.Root()
			total := chain.Total() + 1
			updatedChain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(root).WithHead(minedLink).WithTotal(total).Now()
			if err != nil {
				return err
			}

			// update the chain:
			err = app.application.Sub().Chain().Update(updatedChain)
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
