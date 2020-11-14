package contents

import (
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents/content"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type contents struct {
	lst []content.Content
	mp  map[string]content.Content
}

func createContents(
	lst []content.Content,
	mp map[string]content.Content,
) Contents {
	out := contents{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All return all contents
func (obj *contents) All() []content.Content {
	return obj.lst
}

// Add adds a content to the list
func (obj *contents) Add(content content.Content) error {
	hsh := content.Hash()
	keyname := hsh.String()
	if obj.exists(hsh) {
		str := fmt.Sprintf("the content (hash: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, content)
	obj.mp[keyname] = content
	return nil
}

// NotStored returns the amount of data that is not stored
func (obj *contents) NotStored(file files.File) uint {
	amountNotStored := 0
	chunks := file.Chunks()
	for _, oneChunk := range chunks {
		contentHash := oneChunk.Data()
		if !obj.exists(contentHash) {
			amountNotStored++
		}
	}

	return uint(amountNotStored)
}

func (obj *contents) exists(hsh hash.Hash) bool {
	keyname := hsh.String()
	if _, ok := obj.mp[keyname]; ok {
		return true
	}

	return false
}
