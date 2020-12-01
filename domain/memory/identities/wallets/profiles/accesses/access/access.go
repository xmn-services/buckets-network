package access

import (
	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption/public"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type access struct {
	bucket hash.Hash
	key    public.Key
}

func createAccess(
	bucket hash.Hash,
	key public.Key,
) Access {
	out := access{
		bucket: bucket,
		key:    key,
	}

	return &out
}

// Bucket returns the bucket hash
func (obj *access) Bucket() hash.Hash {
	return obj.bucket
}

// Key returns the public key
func (obj *access) Key() public.Key {
	return obj.key
}
