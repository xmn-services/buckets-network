package windows

// Builder represents the window builder
type Builder interface {
	Create() Builder
	WithWidth(width int) Builder
	WithHeight(height int) Builder
	IsResizable() Builder
	IsFullscreen() Builder
	Now() (Window, error)
}

// Window represents a windows
type Window interface {
	Width() int
	Height() int
	IsResizable() bool
	IsFullscreen() bool
}
