package applications

import (
	"errors"

	"github.com/xmn-services/buckets-network/infrastructure/opengl/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
)

type builder struct {
	nodes nodes.Nodes
}

func createBuilder() Builder {
	out := builder{
		nodes: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithNodes add nodes to the builder
func (app *builder) WithNodes(nodes nodes.Nodes) Builder {
	app.nodes = nodes
	return app
}

// Now builds a new render application instance
func (app *builder) Now() (renders.Application, error) {
	if app.nodes == nil {
		return nil, errors.New("the nodes are mandatory in order to build a render Application instance")
	}

	return createApplication(app.nodes), nil
}
