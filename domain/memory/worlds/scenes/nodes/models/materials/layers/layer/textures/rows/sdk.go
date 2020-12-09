package rows

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows/row"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Builder represents the rows builder
type Builder interface {
	Create() Builder
	WithRows(rows []row.Row) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Rows, error)
}

// Rows represents rows
type Rows interface {
	entities.Immutable
	All() []row.Row
}
