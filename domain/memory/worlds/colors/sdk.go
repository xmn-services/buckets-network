package colors

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a color builder
type Builder interface {
	Create() Builder
	WithRed(red uint32) Builder
	WithGreen(green uint32) Builder
	WithBlue(blue uint32) Builder
	Now() Color
}

// Color represents a color
type Color interface {
	Red() uint32
	Green() uint32
	Blue() uint32
	String() string
}
