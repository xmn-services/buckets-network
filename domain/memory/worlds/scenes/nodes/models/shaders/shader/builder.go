package shader

import (
	"errors"
	"time"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	code             string
	isVertex         bool
	isFragment       bool
	variables []string
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		code:             "",
		isVertex:         false,
		isFragment:       false,
		variables:        nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithCode adds code to the builder
func (app *builder) WithCode(code string) Builder {
	app.code = code
	return app
}

// WithVariables add variables to the builder
func (app *builder) WithVariables(variables []string) Builder {
	app.variables = variables
	return app
}

// IsVertex flags the builder as a vertex shader
func (app *builder) IsVertex() Builder {
	app.isVertex = true
	return app
}

// IsFragment flags the builder as a fragment shader
func (app *builder) IsFragment() Builder {
	app.isFragment = true
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Shader instance
func (app *builder) Now() (Shader, error) {
	if app.variables != nil && len(app.variables) <= 0 {
		app.variables = nil
	}

	if app.variables == nil {
		return nil, errors.New("the variables are mandatory in order to build a Shader instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Shader instance")
	}

	var typ Type
	typeAsString := ""
	if app.isVertex {
		typ = createTypeWithVertex()
		typeAsString = "is_vertex"
	}

	if app.isFragment {
		typ = createTypeWithFragment()
		typeAsString = "is_fragment"
	}

	if typ == nil {
		return nil, errors.New("the Shader must be a Vertex or Fragment type")
	}

	data := [][]byte{
		[]byte(app.code),
		[]byte(typeAsString),
	}

	for _, oneVariable := range app.variables {
		data = append(data, []byte(oneVariable))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	return createShader(immutable, app.code, typ, app.variables), nil
}
