package storages

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/file"
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

// Save saves a file
func (app *application) Save(file file.File) error {
	return nil
}

// Delete deletes a file by hash
func (app *application) Delete(fileHashStr string) error {
	return nil
}
