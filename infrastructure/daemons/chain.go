package daemons

import (
	"time"

	"github.com/xmn-services/buckets-network/application"
	application_chain "github.com/xmn-services/buckets-network/application/chains"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/infrastructure/clients"
)

type chain struct {
	application      application.Application
	remoteAppBuilder clients.Builder
	waitPeriod       time.Duration
	isStarted        bool
}

func createChain(
	application application.Application,
	remoteAppBuilder clients.Builder,
	waitPeriod time.Duration,
) daemons.Application {
	out := chain{
		application:      application,
		remoteAppBuilder: remoteAppBuilder,
		waitPeriod:       waitPeriod,
	}

	return &out
}

// Start starts the application
func (app *chain) Start() error {
	app.isStarted = true

	for {
		// wait period:
		time.Sleep(app.waitPeriod)

		// if the application is not started, continue:
		if !app.isStarted {
			continue
		}

		// retrieve the local chain:
		localChain, err := app.application.Sub().Chain().Retrieve()
		if err != nil {
			return err
		}

		// retrieve the peers:
		localPeers, err := app.application.Sub().Peers().Retrieve()
		if err != nil {
			return err
		}

		biggestDiff := 0
		var biggestChain chains.Chain
		var biggestChainApp application_chain.Application
		allPeers := localPeers.All()
		for _, onePeer := range allPeers {
			remoteApp, err := app.remoteAppBuilder.Create().WithPeer(onePeer).Now()
			if err != nil {
				return err
			}

			remoteChainApp := remoteApp.Sub().Chain()
			remoteChain, err := remoteChainApp.Retrieve()
			if err != nil {
				return err
			}

			diffTrx := int(remoteChain.Total() - localChain.Total())
			if biggestDiff < diffTrx {
				biggestDiff = diffTrx
				biggestChain = remoteChain
				biggestChainApp = remoteChainApp
			}
		}

		// if there is no chain in the network more advanced, continue:
		if biggestChain == nil {
			continue
		}

		// update the chain:
		localIndex := int(localChain.Height())
		diffHeight := int(biggestChain.Height()) - localIndex
		for i := 0; i < diffHeight; i++ {
			chainIndex := localIndex + i
			remoteChainAtIndex, err := biggestChainApp.RetrieveAtIndex(uint(chainIndex))
			if err != nil {
				return err
			}

			remoteHead := remoteChainAtIndex.Head()
			err = app.application.Sub().Chain().Upgrade(remoteHead)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// Stop stops the application
func (app *chain) Stop() error {
	app.isStarted = true
	return nil
}
