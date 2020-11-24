package authenticates

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the authenticate adapter
type Adapter interface {
	Base64ToAuthenticate(encoded string) (Authenticate, error)
	AuthenticateToBase64(auth Authenticate) (string, error)
}

// Builder represents an authenticate builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithSeed(seed string) Builder
	Now() (Authenticate, error)
}

// Authenticate represents an authenticate
type Authenticate interface {
	Name() string
	Password() string
	Seed() string
}
