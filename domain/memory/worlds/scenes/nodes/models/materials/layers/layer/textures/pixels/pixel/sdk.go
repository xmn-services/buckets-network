package pixel

import "github.com/xmn-services/buckets-network/domain/memory/worlds/colors"

// Builder represents a pixel builder
type Builder interface {
	Create() Builder
	WithColor(color colors.Color) Builder
	WithAlpha(alpha uint32) Builder
	Now() (Pixel, error)
}

// Pixel represents a pixel
type Pixel interface {
	Color() colors.Color
	Alpha() uint32
}
