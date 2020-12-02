package accesses

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/accesses/access"

// JSONAccesses represents a JSON access
type JSONAccesses struct {
	List []*access.JSONAccess `json:"access"`
}

func createJSONAccessesFromAccesses(ins Accesses) *JSONAccesses {
	adapter := access.NewAdapter()
	lst := []*access.JSONAccess{}
	accesses := ins.All()
	for _, oneAccess := range accesses {
		lst = append(lst, adapter.ToJSON(oneAccess))
	}

	return createJSONAccesses(
		lst,
	)
}

func createJSONAccesses(
	lst []*access.JSONAccess,
) *JSONAccesses {
	out := JSONAccesses{
		List: lst,
	}

	return &out
}
