package links

import (
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
)

// JSONLink represents a json link
type JSONLink struct {
	PreviousLink string            `json:"previous_link"`
	Next         *mined_blocks.JSONBlock `json:"next"`
	Index        uint              `json:"index"`
	CreatedOn    time.Time         `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *JSONLink {
	previousLink := ins.PreviousLink().String()
	next := mined_blocks.NewAdapter().ToJSON(ins.Next())
	index := ins.Index()
	createdOn := ins.CreatedOn()
	return createJSONLink(previousLink, next, index, createdOn)
}

func createJSONLink(
	previousLink string,
	next *mined_blocks.JSONBlock,
	index uint,
	createdOn time.Time,
) *JSONLink {
	out := JSONLink{
		PreviousLink: previousLink,
		Next:         next,
		Index:        index,
		CreatedOn:    createdOn,
	}

	return &out
}
