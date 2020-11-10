package daemons

import (
	"time"

	application_chain "github.com/xmn-services/buckets-network/application/chains"
	"github.com/xmn-services/buckets-network/application/identities/daemons"
	application_peer "github.com/xmn-services/buckets-network/application/peers"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	client_chain "github.com/xmn-services/buckets-network/infrastructure/clients/chains"
	client_link "github.com/xmn-services/buckets-network/infrastructure/clients/chains/links"
)

type chain struct {
	chainApp              application_chain.Application
	peerApp               application_peer.Application
	remoteChainAppBuilder client_chain.Builder
	remoteLinkAppBuilder  client_link.Builder
	waitPeriod            time.Duration
	isStarted             bool
}

func createChain(
	chainApp application_chain.Application,
	peerApp application_peer.Application,
	remoteChainAppBuilder client_chain.Builder,
	remoteLinkAppBuilder client_link.Builder,
	waitPeriod time.Duration,
) daemons.Application {
	out := chain{
		chainApp:              chainApp,
		peerApp:               peerApp,
		remoteChainAppBuilder: remoteChainAppBuilder,
		remoteLinkAppBuilder:  remoteLinkAppBuilder,
		waitPeriod:            waitPeriod,
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
		localChain, err := app.chainApp.Retrieve()
		if err != nil {
			return err
		}

		// retrieve the peers:
		localPeers, err := app.peerApp.Retrieve()
		if err != nil {
			return err
		}

		biggestDiff := 0
		var biggestChain chains.Chain
		var biggestChainApp application_chain.Application
		allPeers := localPeers.All()
		for _, onePeer := range allPeers {
			remoteChainApp, err := app.remoteChainAppBuilder.Create().WithPeer(onePeer).Now()
			if err != nil {
				return err
			}

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
			err = app.chainApp.Upgrade(remoteHead)
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
