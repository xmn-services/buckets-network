package buckets

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	command_bucket "github.com/xmn-services/buckets-network/application/commands/identities/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
	"github.com/xmn-services/buckets-network/infrastructure/restapis/shared"
)

const baseFormat = "%s%s"

type application struct {
	bucketAdapter buckets.Adapter
	client        *resty.Client
	token         string
	url           string
}

func createApplication(
	bucketAdapter buckets.Adapter,
	client *resty.Client,
	token string,
	peer peer.Peer,
) command_bucket.Application {
	out := application{
		bucketAdapter: bucketAdapter,
		client:        client,
		token:         token,
		url:           fmt.Sprintf(baseFormat, peer.String(), "/identities/buckets"),
	}

	return &out
}

// Add adds a path to the bucket application
func (app *application) Add(absolutePath string) error {
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		SetBody(map[string]string{
			shared.PathKeyname: absolutePath,
		}).
		Post(app.url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// Delete deletes a bucket by hash
func (app *application) Delete(hashStr string) error {
	url := fmt.Sprintf(baseFormat, app.url, hashStr)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Delete(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusOK {
		return nil
	}

	return errors.New(string(resp.Body()))
}

// Retrieve retrieves a bucket by hash
func (app *application) Retrieve(hashStr string) (buckets.Bucket, error) {
	url := fmt.Sprintf(baseFormat, app.url, hashStr)
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return app.bucketAdapter.JSONToBucket(resp.Body())
	}

	return nil, errors.New(string(resp.Body()))
}

// RetrieveAll retrieves all buckets
func (app *application) RetrieveAll() ([]buckets.Bucket, error) {
	resp, err := app.client.R().
		SetHeader(shared.TokenHeadKeyname, app.token).
		Get(app.url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		return app.bucketAdapter.JSONToBuckets(resp.Body())
	}

	return nil, errors.New(string(resp.Body()))
}
