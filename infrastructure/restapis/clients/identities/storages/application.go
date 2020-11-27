package storages

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/file"
)

type application struct {
}

func createApplication() storages.Application {
	out := application{}
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
