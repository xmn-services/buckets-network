package buckets

import (
	"github.com/xmn-services/buckets-network/domain/memory/identities/wallets/miners/buckets/bucket"
)

type builder struct {
	lst []bucket.Bucket
}

func createBuilder() Builder {
	out := builder{
		lst: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBuckets add buckets to the builder
func (app *builder) WithBuckets(buckets []bucket.Bucket) Builder {
	app.lst = buckets
	return app
}

// Now builds a new buckets instance
func (app *builder) Now() (Buckets, error) {
	if app.lst == nil {
		app.lst = []bucket.Bucket{}
	}

	mp := map[string]bucket.Bucket{}
	for _, oneBucket := range app.lst {
		keyname := oneBucket.Hash().String()
		mp[keyname] = oneBucket
	}

	return crateBuckets(app.lst, mp), nil
}
