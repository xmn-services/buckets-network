package chains

import (
	"fmt"

	"github.com/xmn-services/buckets-network/libs/file"
)

type service struct {
	adapter     Adapter
	fileService file.Service
	fileName    string
	extName     string
}

func createService(
	adapter Adapter,
	fileService file.Service,
	fileName string,
	extName string,
) Service {
	out := service{
		adapter:     adapter,
		fileService: fileService,
		fileName:    fileName,
		extName:     extName,
	}

	return &out
}

// Save saves a chain instance
func (app *service) Save(chain Chain) error {
	currentName := fmt.Sprintf("%s.%s", app.fileName, app.extName)
	indexName := fmt.Sprintf("%d.%s", chain.Total()-1, app.extName)
	err := app.save(chain, currentName)
	if err != nil {
		return err
	}

	return app.save(chain, indexName)
}

func (app *service) save(chain Chain, fileName string) error {
	js, err := app.adapter.ToJSON(chain)
	if err != nil {
		return err
	}

	return app.fileService.Save(fileName, js)
}
