package contacts

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts/buckets"
	application_contact_bucket "github.com/xmn-services/buckets-network/application/commands/identities/lists/contacts/buckets"
	"github.com/xmn-services/buckets-network/domain/memory/identities"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts"
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/lists/list/contacts/contact"
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(
	identityRepository identities.Repository,
	identityService identities.Service,
	bucketAppBuilder application_contact_bucket.Builder,
) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
		identityRepository,
		identityService,
		bucketAppBuilder,
	)
}

// NewUpdateBuilder creates a new update builder
func NewUpdateBuilder() UpdateBuilder {
	return createUpdateBuilder()
}

// Builder represents a contact application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	WithList(listHashStr string) Builder
	Now() (Application, error)
}

// Application represents a contacts application
type Application interface {
	RetrieveAll() (contacts.Contacts, error)
	Retrieve(contactHashStr string) (contact.Contact, error)
	Update(contactHashStr string, update Update) error
	Delete(contactHashStr string) error
	Bucket(contactHashStr string) (buckets.Application, error)
}

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithKey(key public.Key) UpdateBuilder
	WithName(name string) UpdateBuilder
	WithDescription(description string) UpdateBuilder
	Now() (Update, error)
}

// Update represents a contact update instance
type Update interface {
	HasKey() bool
	Key() public.Key
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
