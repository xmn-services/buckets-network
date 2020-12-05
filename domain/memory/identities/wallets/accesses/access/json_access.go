package access

import "github.com/xmn-services/buckets-network/libs/cryptography/pk/encryption"

// JSONAccess represents a JSON access
type JSONAccess struct {
	Bucket string `json:"bucket"`
	Key    string `json:"pk"`
}

func createJSONAccessFromAccess(access Access) *JSONAccess {
	privKeyadapter := encryption.NewAdapter()
	key := privKeyadapter.ToEncoded(access.Key())
	bucket := access.Bucket().String()
	return createJSONAccess(
		key,
		bucket,
	)
}

func createJSONAccess(
	key string,
	bucket string,
) *JSONAccess {
	out := JSONAccess{
		Bucket: bucket,
		Key:    key,
	}

	return &out
}
