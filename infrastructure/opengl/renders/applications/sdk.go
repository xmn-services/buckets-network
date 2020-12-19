package applications

import (
	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a renders application builder
type Builder interface {
	Create() Builder
	WithNodes(nodes nodes.Nodes) Builder
	Now() (renders.Application, error)
}
