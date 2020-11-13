package files

import (
	"hash"

	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// File represents a stored file
type File interface {
	entities.Immutable
	File() files.File
	Data() []Data
}

// Data represents a file's data
type Data interface {
	Hash() hash.Hash
	Content() []byte
}
