package links

import (
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
)

// JSONLink represents a json link
type JSONLink struct {
	Previous  string                  `json:"previous"`
	Next      *mined_blocks.JSONBlock `json:"next"`
	Index     uint                    `json:"index"`
	CreatedOn time.Time               `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *JSONLink {
	prev := ins.Previous().String()
	next := mined_blocks.NewAdapter().ToJSON(ins.Next())
	index := ins.Index()
	createdOn := ins.CreatedOn()
	return createJSONLink(prev, next, index, createdOn)
}

func createJSONLink(
	prev string,
	next *mined_blocks.JSONBlock,
	index uint,
	createdOn time.Time,
) *JSONLink {
	out := JSONLink{
		Previous:  prev,
		Next:      next,
		Index:     index,
		CreatedOn: createdOn,
	}

	return &out
}
