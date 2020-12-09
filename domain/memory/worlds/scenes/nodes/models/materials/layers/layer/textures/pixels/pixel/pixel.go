package pixel

import "github.com/xmn-services/buckets-network/domain/memory/worlds/colors"

type pixel struct {
	color colors.Color
	alpha uint32
}

func createPixel(
	color colors.Color,
	alpha uint32,
) Pixel {
	out := pixel{
		color: color,
		alpha: alpha,
	}

	return &out
}

// Color returns the color
func (obj *pixel) Color() colors.Color {
	return obj.color
}

// Alpha returns the alpha
func (obj *pixel) Alpha() uint32 {
	return obj.alpha
}
