package shader

// Builder represents the shader builder
type Builder interface {
	Create() Builder
	WithCode(code string) Builder
	WithVariables(variables []string) Builder
	IsVertex() Builder
	IsFragment() Builder
	Now() (Shader, error)
}

// Shader represents a shader
type Shader interface {
	Code() string
	Type() Type
	Variables() []string
}

// Type represents a shader type
type Type interface {
	IsVertex() bool
	IsFragment() bool
}
