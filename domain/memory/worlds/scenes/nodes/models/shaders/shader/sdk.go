package shader

import (
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the shader builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithCode(code string) Builder
	WithVariables(variables []string) Builder
	IsVertex() Builder
	IsFragment() Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	entities.Immutable
	Index() uint
	Code() string
	Type() Type
	Variables() []string
}

// Type represents a shader type
type Type interface {
	IsVertex() bool
	IsFragment() bool
}
