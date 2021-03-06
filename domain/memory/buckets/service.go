package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
)

type service struct {
	adapter     Adapter
	repository  Repository
	fileService files.Service
	trService   transfer_bucket.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	fileService files.Service,
	trService transfer_bucket.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		fileService: fileService,
		trService:   trService,
	}

	return &out
}

// Save saves an bucket instance
func (app *service) Save(bucket Bucket) error {
	hash := bucket.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	files := bucket.Files()
	err = app.fileService.SaveAll(files)
	if err != nil {
		return err
	}

	trBucket, err := app.adapter.ToTransfer(bucket)
	if err != nil {
		return err
	}

	return app.trService.Save(trBucket)
}

// SaveAll saves buckets
func (app *service) SaveAll(buckets []Bucket) error {
	for _, oneBucket := range buckets {
		err := app.Save(oneBucket)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes an bucket instance
func (app *service) Delete(bucket Bucket) error {
	files := bucket.Files()
	err := app.fileService.DeleteAll(files)
	if err != nil {
		return err
	}

	trBucket, err := app.adapter.ToTransfer(bucket)
	if err != nil {
		return err
	}

	return app.trService.Delete(trBucket)
}
