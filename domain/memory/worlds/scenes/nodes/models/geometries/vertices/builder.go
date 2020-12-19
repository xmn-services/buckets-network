package vertices

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/geometries/vertices/vertex"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter    hash.Adapter
	mutableBuilder entities.MutableBuilder
	hash           *hash.Hash
	withoutHash    bool
	vertices       []vertex.Vertex
	isTriangle     bool
	createdOn      *time.Time
	lastUpdatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	mutableBuilder entities.MutableBuilder,
) Builder {
	out := builder{
		hashAdapter:    hashAdapter,
		mutableBuilder: mutableBuilder,
		hash:           nil,
		withoutHash:    false,
		vertices:       nil,
		isTriangle:     false,
		createdOn:      nil,
		lastUpdatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.mutableBuilder)
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = &hash
	return app
}

// WithoutHash flags the builder as without hash
func (app *builder) WithoutHash() Builder {
	app.withoutHash = true
	return app
}

// WithVertices add the vertices to the builder
func (app *builder) WithVertices(vertices []vertex.Vertex) Builder {
	app.vertices = vertices
	return app
}

// IsTriangle flags the builder as triangles
func (app *builder) IsTriangle() Builder {
	app.isTriangle = true
	return app
}

// CreatedOn adds the creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn adds the lastUpdatedOn time to the builder
func (app *builder) LastUpdatedOn(lastUpdatedOn time.Time) Builder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds the new Vertices instance
func (app *builder) Now() (Vertices, error) {
	if app.vertices == nil {
		app.vertices = []vertex.Vertex{}
	}

	if app.withoutHash {
		typeStr := ""
		if app.isTriangle {
			typeStr = "triangle"
		}

		data := [][]byte{
			[]byte(typeStr),
			[]byte(strconv.Itoa(int(time.Now().UTC().Nanosecond()))),
		}

		for _, oneVertice := range app.vertices {
			data = append(data, []byte(oneVertice.String()))
		}

		hsh, err := app.hashAdapter.FromMultiBytes(data)
		if err != nil {
			return nil, err
		}

		app.hash = hsh
	}

	var typ Type
	if app.isTriangle {
		typ = createTypeWithTriangle()
	}

	if typ == nil {
		return nil, errors.New("the type (isTriangle) is mandatory in order to build a Vertices instance")
	}

	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Vertices instance")
	}

	mutable, err := app.mutableBuilder.Create().WithHash(*app.hash).CreatedOn(app.createdOn).LastUpdatedOn(app.lastUpdatedOn).Now()
	if err != nil {
		return nil, err
	}

	return createVertices(mutable, app.vertices, typ), nil
}
