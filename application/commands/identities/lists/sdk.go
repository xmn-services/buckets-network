package lists

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	identityService identities.Service,
	contactsAppBuilder contacts.Builder,
) Builder {
	hashAdapter := hash.NewAdapter()
	listBuilder := list.NewBuilder()
	return createBuilder(
		hashAdapter,
		listBuilder,
		identityRepository,
		identityService,
		contactsAppBuilder,
	)
}

// Builder represents a list application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents a list application
type Application interface {
	RetrieveAll() (lists.Lists, error)
	Retrieve(listHashStr string) (list.List, error)
	New(name string, description string) error
	Update(listHashStr string, update Update) error
	Delete(listHashStr string) error
	Contacts(listHashStr string) (contacts.Application, error)
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithName(name string) UpdateBuilder
	WithDescription(description string) UpdateBuilder
	Now() (Update, error)
}

// Update represents a list update
type Update interface {
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
