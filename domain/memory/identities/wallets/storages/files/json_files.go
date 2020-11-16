package files

// JSONFiles represents a JSON buckets instance
type JSONFiles struct {
	Files []string `json:"files"`
}

func createJSONFilesFromFiles(ins Files) *JSONFiles {
	files := ins.All()
	lst := []string{}
	for _, oneHash := range files {
		lst = append(lst, oneHash.String())
	}

	return createJSONFiles(lst)
}

func createJSONFiles(
	files []string,
) *JSONFiles {
	out := JSONFiles{
		Files: files,
	}

	return &out
}
