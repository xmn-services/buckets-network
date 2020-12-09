package colors

import "fmt"

type color struct {
	red   uint32
	green uint32
	blue  uint32
}

func createColor(
	red uint32,
	green uint32,
	blue uint32,
) Color {
	out := color{
		red:   red,
		green: green,
		blue:  blue,
	}

	return &out
}

// Red returns the red value
func (obj *color) Red() uint32 {
	return obj.red
}

// Green returns the green value
func (obj *color) Green() uint32 {
	return obj.green
}

// Blue returns the blue value
func (obj *color) Blue() uint32 {
	return obj.blue
}

// String returns the color as string
func (obj *color) String() string {
	return fmt.Sprintf("%d,%d,%d", obj.red, obj.green, obj.blue)
}
