package blocks

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets"
	transfer_block "github.com/xmn-services/buckets-network/domain/transfers/blocks"
)

type service struct {
	adapter       Adapter
	repository    Repository
	bucketService buckets.Service
	trService     transfer_block.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	bucketService buckets.Service,
	trService transfer_block.Service,
) Service {
	out := service{
		adapter:       adapter,
		repository:    repository,
		bucketService: bucketService,
		trService:     trService,
	}

	return &out
}

// Save saves a block
func (app *service) Save(block Block) error {
	hash := block.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	buckets := block.Buckets()
	err = app.bucketService.SaveAll(buckets)
	if err != nil {
		return err
	}

	trBlock, err := app.adapter.ToTransfer(block)
	if err != nil {
		return err
	}

	return app.trService.Save(trBlock)
}
