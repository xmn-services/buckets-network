package files

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files/contents"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type files struct {
	contents []contents.Content
	mp       map[string]contents.Content
}

func createFilesFromJSON(ins *JSONFiles) (Files, error) {
	lst := []contents.Content{}
	contentAdapter := contents.NewAdapter()
	for _, oneContent := range ins.Contents {
		content, err := contentAdapter.ToContent(oneContent)
		if err != nil {
			return nil, err
		}

		lst = append(lst, content)
	}

	return NewBuilder().
		Create().
		WithContents(lst).
		Now()
}

func createFiles(
	contents []contents.Content,
	mp map[string]contents.Content,
) Files {
	out := files{
		contents: contents,
		mp:       mp,
	}

	return &out
}

// All return all contents
func (obj *files) All() []contents.Content {
	return obj.contents
}

// Exists returns true if the chunk hash already exists, false otherwise
func (obj *files) Exists(chunkHash hash.Hash) bool {
	keyname := chunkHash.String()
	if _, ok := obj.mp[keyname]; ok {
		return true
	}

	return false
}

// Add adds a content to the list
func (obj *files) Add(content contents.Content) error {
	chunkHash := content.Chunk()
	keyname := chunkHash.String()
	if obj.Exists(chunkHash) {
		str := fmt.Sprintf("the chunk (hash: %s) already exists and therefore cannot be added", keyname)
		return errors.New(str)
	}

	obj.contents = append(obj.contents, content)
	obj.mp[keyname] = content
	return nil
}

// Delete deletes a chunk hash to the list
func (obj *files) Delete(chunkHash hash.Hash) error {
	keyname := chunkHash.String()
	if !obj.Exists(chunkHash) {
		str := fmt.Sprintf("the file (hash: %s) do not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneContent := range obj.contents {
		if oneContent.Chunk().Compare(chunkHash) {
			obj.contents = append(obj.contents[:index], obj.contents[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *files) MarshalJSON() ([]byte, error) {
	ins := createJSONFilesFromFiles(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *files) UnmarshalJSON(data []byte) error {
	ins := new(JSONFiles)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createFilesFromJSON(ins)
	if err != nil {
		return err
	}

	insFiles := pr.(*files)
	obj.contents = insFiles.contents
	obj.mp = insFiles.mp
	return nil
}
