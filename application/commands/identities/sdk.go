package identities

import (
	"net/url"

	"github.com/xmn-services/buckets-network/application/commands/identities/access"
	"github.com/xmn-services/buckets-network/application/commands/identities/chains"
	"github.com/xmn-services/buckets-network/application/commands/identities/lists"
	"github.com/xmn-services/buckets-network/application/commands/identities/miners"
	"github.com/xmn-services/buckets-network/application/commands/identities/storages"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	minerApp miners.Application,
	accesBuilder access.Builder,
	listBuilder lists.Builder,
	storageBuilder storages.Builder,
	chainBuilder chains.Builder,
	identityRepository identities.Repository,
	identityService identities.Service,
) Builder {
	return createBuilder(
		minerApp,
		accesBuilder,
		listBuilder,
		storageBuilder,
		chainBuilder,
		identityRepository,
		identityService,
	)
}

// NewUpdateBuilder creates a new update builder instance
func NewUpdateBuilder() UpdateBuilder {
	return createUpdateBuilder()
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Application, error)
}

// Application represents an identity application
type Application interface {
	Current() Current
	Sub() SubApplications
}

// Current represents the current application
type Current interface {
	Update(update Update) error
	Retrieve() (identities.Identity, error)
	Delete() error
}

// SubApplications represents an identity's sub applications
type SubApplications interface {
	Access() access.Application
	List() lists.Application
	Storage() storages.Application
	Chain() chains.Application
	Miner() miners.Application
}

// UpdateAdapter represents an update adapter
type UpdateAdapter interface {
	URLValuesToUpdate(values url.Values) (Update, error)
	UpdateToURLValues(update Update) url.Values
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithSeed(seed string) UpdateBuilder
	WithName(name string) UpdateBuilder
	WithPassword(password string) UpdateBuilder
	WithRoot(root string) UpdateBuilder
	Now() (Update, error)
}

// Update represents an identity update
type Update interface {
	HasSeed() bool
	Seed() string
	HasName() bool
	Name() string
	HasPassword() bool
	Password() string
	HasRoot() bool
	Root() string
}
