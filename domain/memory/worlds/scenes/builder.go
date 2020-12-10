package scenes

import (
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	nodes            []nodes.Node
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		nodes:            nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithNodes add nodes to the builder
func (app *builder) WithNodes(nodes []nodes.Node) Builder {
	app.nodes = nodes
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Scene instance
func (app *builder) Now() (Scene, error) {
	if app.nodes != nil && len(app.nodes) <= 0 {
		app.nodes = nil
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
	}

	for _, oneNode := range app.nodes {
		data = append(data, oneNode.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	if app.nodes != nil {
		return createSceneWithNodes(immutable, app.nodes), nil
	}

	return createScene(immutable), nil
}
