package renders

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/renders/render"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the renders builder
type Builder interface {
	Create() Builder
	WithRenders(renders []render.Render) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Renders, error)
}

// Renders represents renders
type Renders interface {
	entities.Immutable
	All() []render.Render
}
