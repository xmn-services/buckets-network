package storages

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
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
) storages.Application {
	out := application{
		client: client,
		token:  token,
		url:    fmt.Sprintf(baseFormat, peer.String(), "/identities/storages"),
	}

	return &out
}

// Save saves a chunk
func (app *application) Save(bucketHashStr string, chunk []byte) error {
	return nil
}

// Delete deletes a chunk
func (app *application) Delete(bucketHashStr string, chunkHashStr string) error {
	return nil
}

// DeleteAll deletes all chunks contained in a bucket
func (app *application) DeleteAll(bucketHashStr string) error {
	return nil
}
