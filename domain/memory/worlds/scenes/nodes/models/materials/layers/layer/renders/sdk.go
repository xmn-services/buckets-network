package renders

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"

// Builder represents the renders builder
type Builder interface {
	Create() Builder
	WithRenders(renders []render.Render) Builder
	Now() (Renders, error)
}

// Renders represents renders
type Renders interface {
	All() []render.Render
}
