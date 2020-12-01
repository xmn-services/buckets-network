package lists

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/profiles/lists/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/profiles/lists/list"
)

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
	Current() Current
	Sub() SubApplications
}

// Current represents a current list application
type Current interface {
	RetrieveAll() lists.Lists
	Retrieve(listHashStr string) (list.List, error)
	New(name string, description string) error
	Update(listHashStr string, update Update) error
	Delete(listHashStr string) error
}

// SubApplications represents a lists sub applications
type SubApplications interface {
	Contact() contacts.Application
}

// Update represents a list update
type Update interface {
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
