package mined

import (
	"github.com/xmn-services/buckets-network/domain/memory/links"
	transfer_mined_link "github.com/xmn-services/buckets-network/domain/transfers/links/mined"
)

type service struct {
	adapter     Adapter
	repository  Repository
	linkService links.Service
	trService   transfer_mined_link.Service
}

func createService(
	adapter Adapter,
	repository Repository,
	linkService links.Service,
	trService transfer_mined_link.Service,
) Service {
	out := service{
		adapter:     adapter,
		repository:  repository,
		linkService: linkService,
		trService:   trService,
	}

	return &out
}

// Save saves a mined link
func (app *service) Save(link Link) error {
	hash := link.Hash()
	_, err := app.repository.Retrieve(hash)
	if err == nil {
		return nil
	}

	lnk := link.Link()
	err = app.linkService.Save(lnk)
	if err != nil {
		return err
	}

	trLink, err := app.adapter.ToTransfer(link)
	if err != nil {
		return err
	}

	return app.trService.Save(trLink)
}
