package links

import (
	transfer_link "github.com/xmn-services/buckets-network/domain/transfers/links"
)

type adapter struct {
	trBuilder transfer_link.Builder
}

func createAdapter(
	trBuilder transfer_link.Builder,
) Adapter {
	out := adapter{
		trBuilder: trBuilder,
	}

	return &out
}

// ToTransfer converts a link to a transfer link instance
func (app *adapter) ToTransfer(link Link) (transfer_link.Link, error) {
	hsh := link.Hash()
	prev := link.Previous()
	next := link.Next().Hash()
	index := link.Index()
	createdOn := link.CreatedOn()
	return app.trBuilder.Create().
		WithHash(hsh).
		WithPrevious(prev).
		WithNext(next).
		WithIndex(index).
		CreatedOn(createdOn).
		Now()
}

// ToJSON converts a link to a JSON link
func (app *adapter) ToJSON(link Link) *JSONLink {
	return createJSONLinkFromLink(link)
}

// ToLink converts a JSON link to a link
func (app *adapter) ToLink(ins *JSONLink) (Link, error) {
	return createLinkFromJSON(ins)
}
