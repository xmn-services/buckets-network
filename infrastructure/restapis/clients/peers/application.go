package peers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	commands_peers "github.com/xmn-services/buckets-network/application/commands/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

const baseFormat = "%s%s"

type application struct {
	peersAdapter peers.Adapter
	peerAdapter  peer.Adapter
	peerBuilder  peer.Builder
	client       *resty.Client
	url          string
}

func createApplication(
	peersAdapter peers.Adapter,
	peerAdapter peer.Adapter,
	peerBuilder peer.Builder,
	client *resty.Client,
	peer peer.Peer,
) commands_peers.Application {
	out := application{
		peersAdapter: peersAdapter,
		peerAdapter:  peerAdapter,
		peerBuilder:  peerBuilder,
		client:       client,
		url:          fmt.Sprintf(baseFormat, peer.String(), "/peers"),
	}

	return &out
}

// Retrieve retrieves a peers instance
func (app *application) Retrieve() (peers.Peers, error) {
	resp, err := app.client.R().
		Get(app.url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		js := resp.Body()
		return app.peersAdapter.JSONToPeers(js)
	}

	return nil, errors.New(string(resp.Body()))
}

// Save saves a peer instance
func (app *application) Save(peer peer.Peer) error {
	urlValues := app.peerAdapter.PeerToURLValues(peer)
	resp, err := app.client.R().
		SetBody(urlValues).
		Post(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// SaveClear saves a clear peer instance
func (app *application) SaveClear(host string, port uint) error {
	peer, err := app.peerBuilder.Create().WithHost(host).WithPort(port).IsClear().Now()
	if err != nil {
		return err
	}

	return app.Save(peer)
}

// SaveOnion saves an onion peer instance
func (app *application) SaveOnion(host string, port uint) error {
	peer, err := app.peerBuilder.Create().WithHost(host).WithPort(port).IsOnion().Now()
	if err != nil {
		return err
	}

	return app.Save(peer)
}
