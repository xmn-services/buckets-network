package identities

import (
	command_identity "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
)

type current struct {
}

func createCurrent() command_identity.Current {
	out := current{}
	return &out
}

// Update updates an identity instance
func (obj *current) Update(update command_identity.Update) error {
	return nil
}

// Retrieve retrieves the identity instance
func (obj *current) Retrieve() (identities.Identity, error) {
	return nil, nil
}

// Delete deletes the identity instance
func (obj *current) Delete() error {
	return nil
}
