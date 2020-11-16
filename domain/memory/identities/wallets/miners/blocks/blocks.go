package blocks

import (
	"encoding/json"
	"errors"
	"fmt"

	mined_blocks "github.com/xmn-services/buckets-network/domain/memory/blocks/mined"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type blocks struct {
	lst []mined_blocks.Block
	mp  map[string]mined_blocks.Block
}

func createBlocksFromJSON(ins *JSONBlocks) (Blocks, error) {
	blocks := []mined_blocks.Block{}
	blockAdapter := mined_blocks.NewAdapter()
	for _, oneJS := range ins.Blocks {
		block, err := blockAdapter.ToBlock(oneJS)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return NewBuilder().
		Create().
		WithBlocks(blocks).
		Now()
}

func createBlocks(
	lst []mined_blocks.Block,
	mp map[string]mined_blocks.Block,
) Blocks {
	out := blocks{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All returns the blocks
func (obj *blocks) All() []mined_blocks.Block {
	return obj.lst
}

// Add adds a block
func (obj *blocks) Add(block mined_blocks.Block) error {
	keyname := block.Hash().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the block (hash: %s) cannot be added because it already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, block)
	obj.mp[keyname] = block
	return nil
}

// Delete deletes a block by hash
func (obj *blocks) Delete(hash hash.Hash) error {
	keyname := hash.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the block (hash: %s) cannot be deleted because it does NOT exists", keyname)
		return errors.New(str)
	}

	for index, oneBlock := range obj.lst {
		if oneBlock.Hash().Compare(hash) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *blocks) MarshalJSON() ([]byte, error) {
	ins := createJSONBlocksFromBlocks(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *blocks) UnmarshalJSON(data []byte) error {
	ins := new(JSONBlocks)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createBlocksFromJSON(ins)
	if err != nil {
		return err
	}

	insBlock := pr.(*blocks)
	obj.lst = insBlock.lst
	obj.mp = insBlock.mp
	return nil
}
