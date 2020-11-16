package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
)

// Factory represents the storages factory
type Factory interface {
	Create() Factory
}

// Storages represents a storages instance
type Storages interface {
	ToDownload() files.Files
	Stored() files.Files
}
