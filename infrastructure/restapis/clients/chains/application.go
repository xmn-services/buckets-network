package chains

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	command_chains "github.com/xmn-services/buckets-network/application/commands/chains"
	"github.com/xmn-services/buckets-network/domain/memory/chains"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

type application struct {
	chainAdapter chains.Adapter
	client       *resty.Client
	url          string
}

func createApplication(
	chainAdapter chains.Adapter,
	client *resty.Client,
	peer peer.Peer,
) command_chains.Application {
	out := application{
		chainAdapter: chainAdapter,
		client:       client,
		url:          fmt.Sprintf("%s%s", peer.String(), "/chains"),
	}

	return &out
}

// Retrieve retrieves a chain application
func (app *application) Retrieve() (chains.Chain, error) {
	resp, err := app.client.R().
		Get(app.url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		js := resp.Body()
		return app.chainAdapter.JSONToChain(js)
	}

	return nil, errors.New(string(resp.Body()))
}

// RetrieveAtIndex retrieves a chain at index application
func (app *application) RetrieveAtIndex(index uint) (chains.Chain, error) {
	url := fmt.Sprintf("%s/%d", app.url, index)
	resp, err := app.client.R().
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		js := resp.Body()
		return app.chainAdapter.JSONToChain(js)
	}

	return nil, errors.New(string(resp.Body()))
}
