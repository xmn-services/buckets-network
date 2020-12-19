package renders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/colors"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/ints"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents a render builder
type Builder interface {
	Create() Builder
	WithOpacity(opacity float64) Builder
	WithViewport(viewport ints.Rectangle) Builder
	WithTexture(tex textures.Texture) Builder
	WithCamera(cam cameras.Camera) Builder
	WithColor(color colors.Color) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Render, error)
}

// Render represents layer render
type Render interface {
	entities.Immutable
	Opacity() float64
	Viewport() ints.Rectangle
	Content() Content
}

// Content represents a render content
type Content interface {
	Hash() hash.Hash
	IsTexture() bool
	Texture() textures.Texture
	IsCamera() bool
	Camera() cameras.Camera
	IsColor() bool
	Color() colors.Color
}
