package contents

import (
	"testing"

	"github.com/xmn-services/buckets-network/libs/hash"
)

func TestFiles_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	bucket, _ := hashAdapter.FromBytes([]byte("bucket"))
	file, _ := hashAdapter.FromBytes([]byte("file"))
	chunk, _ := hashAdapter.FromBytes([]byte("chunk"))

	ins, err := NewBuilder().Create().WithBucket(*bucket).WithFile(*file).WithChunk(*chunk).Now()
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	adapter := NewAdapter()
	jsIns := adapter.ToJSON(ins)
	retIns, err := adapter.ToContent(jsIns)
	if err != nil {
		t.Errorf("the error was expected to nil, error returned: %s", err.Error())
		return
	}

	TestCompare(t, ins, retIns)

}
