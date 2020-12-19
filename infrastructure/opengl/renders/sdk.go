package renders

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/materials"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/textures"
)

// Application represents a render application
type Application interface {
	Render(mat materials.Material) (textures.Texture, error)
}
