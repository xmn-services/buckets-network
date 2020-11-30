package chains

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
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
) chains.Application {
	out := application{
		client: client,
		token:  token,
		url:    fmt.Sprintf(baseFormat, peer.String(), "/identities/chains"),
	}

	return &out
}

// Init initializes a chain application
func (app *application) Init(
	miningValue uint8,
	baseDifficulty uint,
	increasePerBucket float64,
	linkDifficulty uint,
	rootAdditionalBuckets uint,
	headAdditionalBuckets uint,
) error {
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		SetFormDataFromValues(shared.InitChainToURLValues(&shared.InitChain{
			MiningValue:           miningValue,
			BaseDifficulty:        baseDifficulty,
			IncreasePerBucket:     increasePerBucket,
			LinkDifficulty:        linkDifficulty,
			RootAdditionalBuckets: rootAdditionalBuckets,
			HeadAdditionalBuckets: headAdditionalBuckets,
		})).
		Post(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// Block mines a block on the chain
func (app *application) Block(additional uint) error {
	endpoint := fmt.Sprintf("%s%d", "/blocks/", additional)
	url := fmt.Sprintf(baseFormat, app.url, endpoint)

	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Post(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// Link mines a link on the chain
func (app *application) Link(additional uint) error {
	endpoint := fmt.Sprintf("%s%d", "/links/", additional)
	url := fmt.Sprintf(baseFormat, app.url, endpoint)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Post(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}
