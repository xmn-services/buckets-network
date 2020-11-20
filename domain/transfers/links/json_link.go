package links

import (
	"time"
)

type jsonLink struct {
	Hash      string    `json:"hash"`
	Previous  string    `json:"previous"`
	Next      string    `json:"next"`
	Index     uint      `json:"index"`
	CreatedOn time.Time `json:"created_on"`
}

func createJSONLinkFromLink(ins Link) *jsonLink {
	hash := ins.Hash().String()
	prev := ins.Previous().String()
	next := ins.Next().String()
	index := ins.Index()
	createdOn := ins.CreatedOn()
	return createJSONLink(hash, prev, next, index, createdOn)
}

func createJSONLink(
	hash string,
	prev string,
	next string,
	index uint,
	createdOn time.Time,
) *jsonLink {
	out := jsonLink{
		Hash:      hash,
		Previous:  prev,
		Next:      next,
		Index:     index,
		CreatedOn: createdOn,
	}

	return &out
}
