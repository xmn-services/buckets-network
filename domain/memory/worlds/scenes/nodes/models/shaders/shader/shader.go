package shader

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type shader struct {
	immutable entities.Immutable
	code      string
	typ       Type
	variables []string
}

func createShader(
	immutable entities.Immutable,
	code string,
	typ Type,
	variables []string,
) Shader {
	out := shader{
		immutable: immutable,
		code:      code,
		typ:       typ,
		variables: variables,
	}

	return &out
}

// Hash returns the hash
func (obj *shader) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Code returns the code
func (obj *shader) Code() string {
	return obj.code
}

// Type returns the type
func (obj *shader) Type() Type {
	return obj.typ
}

// Variables returns the variables
func (obj *shader) Variables() []string {
	return obj.variables
}

// CreatedOn returns the creation time
func (obj *shader) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}
