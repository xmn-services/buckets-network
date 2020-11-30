package contents

// JSONContent represents a JSON content instance
type JSONContent struct {
	Bucket string `json:"bucket"`
	File   string `json:"file"`
	Chunk  string `json:"chunk"`
}

func createJSONContentFromContent(ins Content) *JSONContent {
	bucket := ins.Bucket().String()
	file := ins.File().String()
	chunk := ins.Chunk().String()
	return createJSONContent(bucket, file, chunk)
}

func createJSONContent(
	bucket string,
	file string,
	chunk string,
) *JSONContent {
	out := JSONContent{
		Bucket: bucket,
		File:   file,
		Chunk:  chunk,
	}

	return &out
}
