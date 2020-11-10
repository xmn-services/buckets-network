package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	peers_mem "github.com/xmn-services/buckets-network/domain/memory/peers"
	client_bucket "github.com/xmn-services/buckets-network/infrastructure/clients/identities/buckets"
)

type follow struct {
	peersRepository        peers_mem.Repository
	identityRepository     identities.Repository
	identityService        identities.Service
	remoteBucketAppBuilder client_bucket.Builder
	waitPeriod             time.Duration
	name                   string
	password               string
	seed                   string
	isStarted              bool
}

func createFollow(
	peersRepository peers_mem.Repository,
	identityRepository identities.Repository,
	identityService identities.Service,
	remoteBucketAppBuilder client_bucket.Builder,
	waitPeriod time.Duration,
	name string,
	password string,
	seed string,
) daemons.Application {
	out := follow{
		peersRepository:        peersRepository,
		identityRepository:     identityRepository,
		identityService:        identityService,
		remoteBucketAppBuilder: remoteBucketAppBuilder,
		waitPeriod:             waitPeriod,
		name:                   name,
		password:               password,
		seed:                   seed,
		isStarted:              false,
	}

	return &out
}

// Start starts the application
func (app *follow) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the identity:
		identity, err := app.identityRepository.Retrieve(app.name, app.password, app.seed)
		if err != nil {
			return err
		}

		// retrieve the peers:
		peers, err := app.peersRepository.Retrieve()
		if err != nil {
			return err
		}

		// build the client bucket:
		removeBucketApp, err := app.remoteBucketAppBuilder.Create().
			WithName(app.name).
			WithPassword(app.password).
			WithSeed(app.seed).
			WithPeers(peers).
			Now()

		if err != nil {
			return err
		}

		followBuckets := identity.Wallet().Follows().Requests()
		for _, oneBucketHash := range followBuckets {
			bucket, err := removeBucketApp.Retrieve(oneBucketHash)
			if err != nil {
				return err
			}

			err = identity.Wallet().Follows().Add(bucket)
			if err != nil {
				return err
			}
		}

		// save the identity:
		err = app.identityService.Update(
			identity.Hash(),
			identity,
			app.password,
			app.password,
		)

		if err != nil {
			return err
		}
	}
}

// Stop stops the application
func (app *follow) Stop() error {
	app.isStarted = true
	return nil
}
