package layers

import (
	"errors"

	domain_layers "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/layers/layer"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type builder struct {
	layerBuilder layer.Builder
	layers       domain_layers.Layers
	prog         programs.Program
}

func createBuilder(
	layerBuilder layer.Builder,
) Builder {
	out := builder{
		layerBuilder: layerBuilder,
		layers:       nil,
		prog:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.layerBuilder)
}

// WithLayers add layers to the builder
func (app *builder) WithLayers(layers domain_layers.Layers) Builder {
	app.layers = layers
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(prog programs.Program) Builder {
	app.prog = prog
	return app
}

// Now builds a new Layers instance
func (app *builder) Now() (Layers, error) {
	if app.layers == nil {
		return nil, errors.New("the layers are mandatory in order to build a Layers instance")
	}

	list := []layer.Layer{}
	all := app.layers.All()
	for _, oneDomainLayer := range all {
		layerBuilder := app.layerBuilder.Create().WithLayer(oneDomainLayer)
		if app.prog != nil {
			layerBuilder.WithProgram(app.prog)
		}

		layer, err := layerBuilder.Now()
		if err != nil {
			return nil, err
		}

		list = append(list, layer)
	}

	return createLayers(list), nil
}
