package shader

import (
	"time"

	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents the shader builder
type Builder interface {
	Create() Builder
	WithCode(code string) Builder
	WithVariables(variables []string) Builder
	IsVertex() Builder
	IsFragment() Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	entities.Immutable
	Code() string
	Type() Type
	Variables() []string
}

// Type represents a shader type
type Type interface {
	IsVertex() bool
	IsFragment() bool
}
