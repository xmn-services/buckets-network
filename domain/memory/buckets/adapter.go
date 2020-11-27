package buckets

import (
	"encoding/json"

	transfer_bucket "github.com/xmn-services/buckets-network/domain/transfers/buckets"
	"github.com/xmn-services/buckets-network/libs/hashtree"
)

type adapter struct {
	hashTreeBuilder hashtree.Builder
	trBuilder       transfer_bucket.Builder
}

func createAdapter(
	hashTreeBuilder hashtree.Builder,
	trBuilder transfer_bucket.Builder,
) Adapter {
	out := adapter{
		hashTreeBuilder: hashTreeBuilder,
		trBuilder:       trBuilder,
	}

	return &out
}

// ToTransfer converts an bucket to a transfer bucket
func (app *adapter) ToTransfer(bucket Bucket) (transfer_bucket.Bucket, error) {
	hash := bucket.Hash()
	files := bucket.Files()

	blocks := [][]byte{}
	for _, oneFile := range files {
		blocks = append(blocks, oneFile.Hash().Bytes())
	}

	ht, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	amount := uint(len(files))
	createdOn := bucket.CreatedOn()
	return app.trBuilder.Create().WithHash(hash).WithFiles(ht).WithAmount(amount).CreatedOn(createdOn).Now()
}

// ToJSON converts bucket to JSON
func (app *adapter) ToJSON(bucket Bucket) *JSONBucket {
	return createJSONBucketFromBucket(bucket)
}

// ToBucket converts JSON to bucket
func (app *adapter) ToBucket(ins *JSONBucket) (Bucket, error) {
	return createBucketFromJSON(ins)
}

// JSONToBucket converts json to Bucket instance
func (app *adapter) JSONToBucket(js []byte) (Bucket, error) {
	ins := new(JSONBucket)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	return app.ToBucket(ins)
}

// JSONToBuckets coverts JSON to buckets instance
func (app *adapter) JSONToBuckets(js []byte) ([]Bucket, error) {
	ins := new([]JSONBucket)
	err := json.Unmarshal(js, ins)
	if err != nil {
		return nil, err
	}

	out := []Bucket{}
	for _, oneIns := range *ins {
		bucket, err := app.ToBucket(&oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, bucket)
	}

	return out, nil
}
