package peers

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/libs/file"
)

type service struct {
	fileService     file.Service
	fileNameWithExt string
}

func createService(
	fileService file.Service,
	fileNameWithExt string,
) Service {
	out := service{
		fileService:     fileService,
		fileNameWithExt: fileNameWithExt,
	}

	return &out
}

// Save save peers
func (app *service) Save(peers Peers) error {
	js, err := json.Marshal(peers)
	if err != nil {
		return err
	}

	return app.fileService.Save(app.fileNameWithExt, js)
}
