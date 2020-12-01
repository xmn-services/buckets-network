package profiles

import (
	"github.com/xmn-services/buckets-network/application/commands/identities/profiles/access"
	"github.com/xmn-services/buckets-network/application/commands/identities/profiles/lists"
)

// Application represents a profile application
type Application interface {
	Current() Current
	Sub() SubApplications
}

// Current represents the current profile application
type Current interface {
	Update(update Update) error
}

// SubApplications represents a profile suv applications
type SubApplications interface {
	Access() access.Application
	List() lists.Application
}

// Update represents an update instance
type Update interface {
	HasName() bool
	Name() string
	HasDescription() bool
	Description() string
}
