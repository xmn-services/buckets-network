package storages

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Factory represents the storages factory
type Factory interface {
	Create() Factory
}

// Storages represents a storages instance
type Storages interface {
	entities.Mutable
	ToDownload() files.Files
	Stored() files.Files
}
