package contacts

import "github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"

type update struct {
	key         public.Key
	name        string
	description string
}

func createUpdateWithKeyAndNameAndDescription(
	key public.Key,
	name string,
	description string,
) Update {
	return createUpdateInternally(key, name, description)
}

func createUpdateWithKeyAndName(
	key public.Key,
	name string,
) Update {
	return createUpdateInternally(key, name, "")
}

func createUpdateWithKeyAndDescription(
	key public.Key,
	description string,
) Update {
	return createUpdateInternally(key, "", description)
}

func createUpdateWithNameAndDescription(
	name string,
	description string,
) Update {
	return createUpdateInternally(nil, name, description)
}

func createUpdateWithKey(
	key public.Key,
) Update {
	return createUpdateInternally(key, "", "")
}

func createUpdateWithName(
	name string,
) Update {
	return createUpdateInternally(nil, name, "")
}

func createUpdateWithDescription(
	description string,
) Update {
	return createUpdateInternally(nil, "", description)
}

func createUpdateInternally(
	key public.Key,
	name string,
	description string,
) Update {
	out := update{
		key:         key,
		name:        name,
		description: description,
	}

	return &out
}

// HasKey returns true if there is a key, false otherwise
func (obj *update) HasKey() bool {
	return obj.key != nil
}

// Key returns the key, if any
func (obj *update) Key() public.Key {
	return obj.key
}

// HasName returns true if there is a name, false otherwise
func (obj *update) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *update) Name() string {
	return obj.name
}

// HasDescription returns true if there is a description, false otherwise
func (obj *update) HasDescription() bool {
	return obj.description != ""
}

// Description returns the description, if any
func (obj *update) Description() string {
	return obj.description
}
