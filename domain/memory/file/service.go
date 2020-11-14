package file

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	transfer_data "github.com/xmn-services/buckets-network/domain/transfers/buckets/files/chunks/data"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type service struct {
	repository    Repository
	fileService   files.Service
	trDataService transfer_data.Service
}

func createService(
	repository Repository,
	fileService files.Service,
	trDataService transfer_data.Service,
) Service {
	out := service{
		repository:    repository,
		fileService:   fileService,
		trDataService: trDataService,
	}

	return &out
}

// Delete deletes a file
func (app *service) Delete(hash hash.Hash) error {
	storedFile, err := app.repository.Retrieve(hash)
	if err != nil {
		return err
	}

	contents := storedFile.Contents().All()
	for _, oneContent := range contents {
		hsh := oneContent.Hash()
		err = app.trDataService.Delete(hsh)
		if err != nil {
			return err
		}
	}

	file := storedFile.File()
	return app.fileService.Delete(file)
}
