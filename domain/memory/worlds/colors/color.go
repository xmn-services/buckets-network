package colors

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
