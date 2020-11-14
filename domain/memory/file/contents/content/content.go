package content

import (
	"github.com/xmn-services/buckets-network/libs/hash"
)

type content struct {
	hash    hash.Hash
	content []byte
}

func createContent(
	hash hash.Hash,
	data []byte,
) Content {
	out := content{
		hash:    hash,
		content: data,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *content) Content() []byte {
	return obj.content
}
