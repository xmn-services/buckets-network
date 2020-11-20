package links

import (
	"encoding/json"
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type link struct {
	immutable entities.Immutable
	prev      hash.Hash
	next      hash.Hash
	index     uint
}

func createLinkFromJSON(ins *jsonLink) (Link, error) {
	hashAdapter := hash.NewAdapter()
	hsh, err := hashAdapter.FromString(ins.Hash)
	if err != nil {
		return nil, err
	}

	prev, err := hashAdapter.FromString(ins.Previous)
	if err != nil {
		return nil, err
	}

	next, err := hashAdapter.FromString(ins.Next)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithHash(*hsh).
		WithPrevious(*prev).
		WithNext(*next).
		WithIndex(ins.Index).
		CreatedOn(ins.CreatedOn).
		Now()
}

func createLink(
	immutable entities.Immutable,
	prev hash.Hash,
	next hash.Hash,
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

// Next returns the next hash
func (obj *link) Next() hash.Hash {
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
	ins := new(jsonLink)
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
