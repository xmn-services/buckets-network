package miners

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

const baseFormat = "%s%s"

type application struct {
	client *resty.Client
	token  string
	url    string
}

func createApplication(
	client *resty.Client,
	token string,
	peer peer.Peer,
) miners.Application {
	out := application{
		client: client,
		token:  token,
		url:    fmt.Sprintf(baseFormat, peer.String(), "/miners"),
	}
	return &out
}

// Test executes a test on the miner application
func (app *application) Test(difficulty uint) (string, error) {
	endpoint := fmt.Sprintf("%s%d", "/test/", difficulty)
	url := fmt.Sprintf(baseFormat, app.url, endpoint)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return string(resp.Body()), nil
	}

	return "", errors.New(string(resp.Body()))
}

// Block mines a block
func (app *application) Block(blockHashStr string) (string, error) {
	endpoint := fmt.Sprintf(baseFormat, "/block/", blockHashStr)
	url := fmt.Sprintf(baseFormat, app.url, endpoint)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return string(resp.Body()), nil
	}

	return "", errors.New(string(resp.Body()))
}

// Link mines a link
func (app *application) Link(linkHashStr string) (string, error) {
	endpoint := fmt.Sprintf(baseFormat, "/link/", linkHashStr)
	url := fmt.Sprintf(baseFormat, app.url, endpoint)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusOK {
		return string(resp.Body()), nil
	}

	return "", errors.New(string(resp.Body()))
}
