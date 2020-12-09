package colors

// Builder represents a color builder
type Builder interface {
	Create() Builder
	WithRed(red uint32) Builder
	WithGreen(green uint32) Builder
	WithBlue(blue uint32) Builder
	Now() (Color, error)
}

// Color represents a color
type Color interface {
	Red() uint32
	Green() uint32
	Blue() uint32
}
