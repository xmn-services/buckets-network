package identities

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	command_identity "github.com/xmn-services/buckets-network/application/commands/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/peers/peer"
)

const baseFormat = "%s%s"

type current struct {
	client *resty.Client
	token  string
	url    string
}

func createCurrent(
	client *resty.Client,
	token string,
	peer peer.Peer,
) command_identity.Current {
	out := current{
		client: client,
		token:  token,
		url:    fmt.Sprintf(baseFormat, peer.String(), "/identities"),
	}

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
