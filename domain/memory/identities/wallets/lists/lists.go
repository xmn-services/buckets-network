package lists

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type lists struct {
	lst []list.List
	mp  map[string]list.List
}

func createListsFromJSON(ins *JSONLists) (Lists, error) {
	lst := []list.List{}
	adapter := list.NewAdapter()
	for _, oneList := range ins.List {
		list, err := adapter.ToList(oneList)
		if err != nil {
			return nil, err
		}

		lst = append(lst, list)
	}

	return NewBuilder().
		Create().
		WithList(lst).
		Now()
}

func createLists(
	lst []list.List,
	mp map[string]list.List,
) Lists {
	out := lists{
		lst: lst,
		mp:  mp,
	}

	return &out
}

// All returns the lists
func (obj *lists) All() []list.List {
	return obj.lst
}

// Add adds a list
func (obj *lists) Add(list list.List) error {
	keyname := list.Hash().String()
	if _, ok := obj.mp[keyname]; ok {
		str := fmt.Sprintf("the list (hash: %s) already exists", keyname)
		return errors.New(str)
	}

	obj.lst = append(obj.lst, list)
	obj.mp[keyname] = list
	return nil
}

// Delete deletes a list
func (obj *lists) Delete(listHash hash.Hash) error {
	keyname := listHash.String()
	if _, ok := obj.mp[keyname]; !ok {
		str := fmt.Sprintf("the list (hash: %s) does not exists and therefore cannot be deleted", keyname)
		return errors.New(str)
	}

	for index, oneHash := range obj.lst {
		if oneHash.Hash().Compare(listHash) {
			obj.lst = append(obj.lst[:index], obj.lst[index+1:]...)
			break
		}
	}

	delete(obj.mp, keyname)
	return nil
}

// MarshalJSON converts the instance to JSON
func (obj *lists) MarshalJSON() ([]byte, error) {
	ins := createJSONListsFromLists(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *lists) UnmarshalJSON(data []byte) error {
	ins := new(JSONLists)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createListsFromJSON(ins)
	if err != nil {
		return err
	}

	insLists := pr.(*lists)
	obj.lst = insLists.lst
	obj.mp = insLists.mp
	return nil
}
