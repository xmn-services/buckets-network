package access

import (
	"encoding/json"

	"github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type access struct {
	bucket hash.Hash
	key    encryption.PrivateKey
}

func createAccessFromJSON(ins *JSONAccess) (Access, error) {
	bucket, err := hash.NewAdapter().FromString(ins.Bucket)
	if err != nil {
		return nil, err
	}

	key, err := encryption.NewAdapter().FromEncoded(ins.Key)
	if err != nil {
		return nil, err
	}

	return NewBuilder().
		Create().
		WithBucket(*bucket).
		WithKey(key).
		Now()
}

func createAccess(
	bucket hash.Hash,
	key encryption.PrivateKey,
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
func (obj *access) Key() encryption.PrivateKey {
	return obj.key
}

// MarshalJSON converts the instance to JSON
func (obj *access) MarshalJSON() ([]byte, error) {
	ins := createJSONAccessFromAccess(obj)
	return json.Marshal(ins)
}

// UnmarshalJSON converts the JSON to an instance
func (obj *access) UnmarshalJSON(data []byte) error {
	ins := new(JSONAccess)
	err := json.Unmarshal(data, ins)
	if err != nil {
		return err
	}

	pr, err := createAccessFromJSON(ins)
	if err != nil {
		return err
	}

	insAccess := pr.(*access)
	obj.bucket = insAccess.bucket
	obj.key = insAccess.key
	return nil
}
