package lists

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"

// JSONLists represents a JSON list
type JSONLists struct {
	List []*list.JSONList `json:"list"`
}

func createJSONListsFromLists(ins Lists) *JSONLists {
	adapter := list.NewAdapter()
	lst := []*list.JSONList{}
	lists := ins.All()
	for _, oneList := range lists {
		lst = append(lst, adapter.ToJSON(oneList))
	}

	return createJSONLists(
		lst,
	)
}

func createJSONLists(
	lst []*list.JSONList,
) *JSONLists {
	out := JSONLists{
		List: lst,
	}

	return &out
}
