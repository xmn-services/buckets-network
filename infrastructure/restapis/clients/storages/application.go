package storages

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	commands_storages "github.com/xmn-services/buckets-network/application/commands/storages"
	stored_file "github.com/xmn-services/buckets-network/domain/memory/file"
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

// IsStored retrieves true if the file is totally stored, false otherwise
func (app *application) IsStored(fileHashStr string) bool {
	return true
}

// Retrieve retrieves a stored file instance by hash
func (app *application) Retrieve(fileHashStr string) (stored_file.File, error) {
	return nil, nil
}
