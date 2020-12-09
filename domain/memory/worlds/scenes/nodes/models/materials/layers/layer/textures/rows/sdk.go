package rows

import "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures/rows/row"

// Builder represents the rows builder
type Builder interface {
	Create() Builder
	WithRows(rows []row.Row) Builder
	Now() (Rows, error)
}

// Rows represents rows
type Rows interface {
	All() []row.Row
}
