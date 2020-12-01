package profiles

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/accesses"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists"
)

type profile struct {
	name        string
	description string
	list        lists.Lists
	access      accesses.Accesses
}

func createProfile(
	name string,
	description string,
	list lists.Lists,
	access accesses.Accesses,
) Profile {
	out := profile{
		name:        name,
		description: description,
		list:        list,
		access:      access,
	}

	return &out
}

// Name returns the name
func (obj *profile) Name() string {
	return obj.name
}

// Description returns the description
func (obj *profile) Description() string {
	return obj.description
}

// List returns the list
func (obj *profile) List() lists.Lists {
	return obj.list
}

// Access returns the access
func (obj *profile) Access() accesses.Accesses {
	return obj.access
}
