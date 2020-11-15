package buckets

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts Buckets instance to JSON
func (app *adapter) ToJSON(ins Buckets) *JSONBuckets {
	return createJSONBucketsFromBuckets(ins)
}

// ToBuckets creates a JSON to Buckets instance
func (app *adapter) ToBuckets(js *JSONBuckets) (Buckets, error) {
	return createBucketsFromJSON(js)
}
