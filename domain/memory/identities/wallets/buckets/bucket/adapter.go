package bucket

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts bucket instance to JSON
func (app *adapter) ToJSON(ins Bucket) *JSONBucket {
	return createJSONBucketFromBucket(ins)
}

// ToBucket converts json to bucket instance
func (app *adapter) ToBucket(js *JSONBucket) (Bucket, error) {
	return createBucketFromJSON(js)
}
