package files

import "github.com/xmn-services/buckets-network/domain/memory/identities/wallets/storages/files/contents"

// JSONFiles represents a JSON buckets instance
type JSONFiles struct {
	Contents []*contents.JSONContent `json:"contents"`
}

func createJSONFilesFromFiles(ins Files) *JSONFiles {
	lst := ins.All()
	out := []*contents.JSONContent{}
	contentAdapter := contents.NewAdapter()
	for _, oneContent := range lst {
		content := contentAdapter.ToJSON(oneContent)
		out = append(out, content)
	}

	return createJSONFiles(out)
}

func createJSONFiles(
	contents []*contents.JSONContent,
) *JSONFiles {
	out := JSONFiles{
		Contents: contents,
	}

	return &out
}
