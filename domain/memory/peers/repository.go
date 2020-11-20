package peers

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/libs/file"
)

type repository struct {
	fileRepository  file.Repository
	fileNameWithExt string
}

func createRepository(
	fileRepository file.Repository,
	fileNameWithExt string,
) Repository {
	out := repository{
		fileRepository:  fileRepository,
		fileNameWithExt: fileNameWithExt,
	}

	return &out
}

// Exists returns true if there is peers, false otherwise
func (app *repository) Exists() bool {
	return app.fileRepository.Exists(app.fileNameWithExt)
}

// Retrieve retrieve peers
func (app *repository) Retrieve() (Peers, error) {
	js, err := app.fileRepository.Retrieve(app.fileNameWithExt)
	if err != nil {
		return nil, err
	}

	ins := new(peers)
	err = json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
