package chains

import (
	"fmt"

	"github.com/xmn-services/buckets-network/libs/file"
)

type repository struct {
	adapter        Adapter
	fileRepository file.Repository
	fileName       string
	extName        string
}

func createRepository(
	adapter Adapter,
	fileRepository file.Repository,
	fileName string,
	extName string,
) Repository {
	out := repository{
		adapter:        adapter,
		fileRepository: fileRepository,
		fileName:       fileName,
		extName:        extName,
	}

	return &out
}

// Retrieve retrieves a chain by hash
func (app *repository) Retrieve() (Chain, error) {
	name := fmt.Sprintf("%s.%s", app.fileName, app.extName)
	return app.retrieveByFileName(name)
}

// RetrieveAtIndex retrieves a chain at index
func (app *repository) RetrieveAtIndex(index uint) (Chain, error) {
	name := fmt.Sprintf("%d.%s", index, app.extName)
	return app.retrieveByFileName(name)
}

func (app *repository) retrieveByFileName(name string) (Chain, error) {
	js, err := app.fileRepository.Retrieve(name)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToChain(js)
}
