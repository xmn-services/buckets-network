package nodes

import (
	"errors"

	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
)

type builder struct {
	nodeBuilder NodeBuilder
	list        []domain_nodes.Node
}

func createBuilder(
	nodeBuilder NodeBuilder,
) Builder {
	out := builder{
		nodeBuilder: nodeBuilder,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.nodeBuilder)
}

// WithNodes add nodes to the builder
func (app *builder) WithNodes(nodes []domain_nodes.Node) Builder {
	app.list = nodes
	return app
}

// Now builds a new Nodes instance
func (app *builder) Now() (Nodes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Node in order to build a Nodes instance")
	}

	list := []Node{}
	for _, oneDomainNode := range app.list {
		node, err := app.nodeBuilder.Create().WithNode(oneDomainNode).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, node)
	}

	return createNodes(list), nil
}
