package file

import (
	"github.com/xmn-services/buckets-network/domain/memory/buckets/files"
	"github.com/xmn-services/buckets-network/domain/memory/file/contents"
)

type file struct {
	file     files.File
	contents contents.Contents
}

func createFile(
	fil files.File,
	contents contents.Contents,
) File {
	out := file{
		file:     fil,
		contents: contents,
	}

	return &out
}

// File returns the file
func (obj *file) File() files.File {
	return obj.file
}

// Contents returns the contents
func (obj *file) Contents() contents.Contents {
	return obj.contents
}
