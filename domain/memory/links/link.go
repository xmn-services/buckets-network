package links

import (
	"encoding/json"
	"time"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type link struct {
	immutable entities.Immutable
	prev      hash.Hash
	next      mined_blocks.Block
	index     uint
}

func createLinkFromJSON(ins *JSONLink) (Link, error) {
	hashAdapter := hash.NewAdapter()
	prev, err := hashAdapter.FromString(ins.Previous)
	if err != nil {
		return nil, err
	}

	blocksAdapter := mined_blocks.NewAdapter()
	next, err := blocksAdapter.ToBlock(ins.Next)
	if err != nil {
		return nil, err
	}

	return NewBuilder().Create().
		WithIndex(ins.Index).
		WithNext(next).
		WithPrevious(*prev).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLink(
	immutable entities.Immutable,
	prev hash.Hash,
	next mined_blocks.Block,
	index uint,
) Link {
	out := link{
		immutable: immutable,
		prev:      prev,
		next:      next,
		index:     index,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Previous returns the previous hash
func (obj *link) Previous() hash.Hash {
	return obj.prev
}

// Next returns the next block
func (obj *link) Next() mined_blocks.Block {
	return obj.next
}

// Index returns the index
func (obj *link) Index() uint {
	return obj.index
}

// CreatedOn returns the creation time
func (obj *link) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// MarshalJSON converts the instance to JSON
func (obj *link) MarshalJSON() ([]byte, error) {
	ins := createJSONLinkFromLink(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *link) UnmarshalJSON(data []byte) error {
	ins := new(JSONLink)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createLinkFromJSON(ins)
	if err != nil {
		return err
	}

	insLink := pr.(*link)
	obj.immutable = insLink.immutable
	obj.prev = insLink.prev
	obj.next = insLink.next
	obj.index = insLink.index
	return nil
}
