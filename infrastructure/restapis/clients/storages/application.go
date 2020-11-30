package storages

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	commands_storages "github.com/xmn-services/buckets-network/application/commands/storages"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

const baseFormat = "%s%s"

type application struct {
	client *resty.Client
	url    string
}

func createApplication(
	client *resty.Client,
	peer peer.Peer,
) commands_storages.Application {
	out := application{
		client: client,
		url:    fmt.Sprintf(baseFormat, peer.String(), "/storages"),
	}

	return &out
}

// Exists returns true if the chunk exists, false otherwise
func (app *application) Exists(bucketHashStr string, chunkHashStr string) bool {
	return true
}

// Retrieve retrieves a chunk from a bucket
func (app *application) Retrieve(bucketHashStr string, chunkHashStr string) ([]byte, error) {
	return nil, nil
}
