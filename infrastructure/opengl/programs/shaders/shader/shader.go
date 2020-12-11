package shader

import "github.com/xmn-services/buckets-network/libs/hash"

type shader struct {
	sh hash.Hash
	id uint32
}

func createShader(
	sh hash.Hash,
	id uint32,
) Shader {
	out := shader{
		sh: sh,
		id: id,
	}

	return &out
}

// Shader returns the shader
func (obj *shader) Shader() hash.Hash {
	return obj.sh
}

// Identifier returns the identifier
func (obj *shader) Identifier() uint32 {
	return obj.id
}
