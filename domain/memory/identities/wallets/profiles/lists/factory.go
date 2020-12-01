package lists

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list"

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new lists instance
func (app *factory) Create() Lists {
	lst := []list.List{}
	mp := map[string]list.List{}
	return createLists(lst, mp)
}
