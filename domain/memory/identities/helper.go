package identities

import (
	"fmt"

	"github.com/xmn-services/buckets-network/libs/hash"
)

func makePassword(hashAdapter hash.Adapter, seed string, password string) (string, error) {
	hsh, err := hashAdapter.FromMultiBytes([][]byte{
		[]byte(password),
		[]byte(seed),
	})

	if err != nil {
		return "", err
	}

	return hsh.String(), nil
}

func makeFileName(name string, extension string) string {
	return fmt.Sprintf("%s.%s", name, extension)
}
