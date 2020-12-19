package cameras

import (
	"errors"
	"strconv"
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder entities.ImmutableBuilder
	lookAtVariable   string
	eye              *fl32.Vec3
	center           *fl32.Vec3
	up               *fl32.Vec3
	projVariable     string
	fov              *float32
	aspectRatio      *float32
	near             *float32
	far              *float32
	index            uint
	createdOn        *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder entities.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		lookAtVariable:   "",
		eye:              nil,
		center:           nil,
		up:               nil,
		projVariable:     "",
		fov:              nil,
		aspectRatio:      nil,
		near:             nil,
		far:              nil,
		index:            uint(0),
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.immutableBuilder,
	)
}

// WithLookAtVariable adds the lookAt variable to the builder
func (app *builder) WithLookAtVariable(lookAtVariable string) Builder {
	app.lookAtVariable = lookAtVariable
	return app
}

// WithLookAtEye adds the lookAt eye to the builder
func (app *builder) WithLookAtEye(eye fl32.Vec3) Builder {
	app.eye = &eye
	return app
}

// WithLookAtCenter adds the lookAt center to the builder
func (app *builder) WithLookAtCenter(center fl32.Vec3) Builder {
	app.center = &center
	return app
}

// WithLookAtUp adds the lookAt up to the builder
func (app *builder) WithLookAtUp(up fl32.Vec3) Builder {
	app.up = &up
	return app
}

// WithProjectionVariable adds the projection variable to the builder
func (app *builder) WithProjectionVariable(projVariable string) Builder {
	app.projVariable = projVariable
	return app
}

// WithProjectionFieldofView adds the projection fov to the builder
func (app *builder) WithProjectionFieldofView(fov float32) Builder {
	app.fov = &fov
	return app
}

// WithProjectionAspectRatio adds the projection aspectRatio to the builder
func (app *builder) WithProjectionAspectRatio(aspectRatio float32) Builder {
	app.aspectRatio = &aspectRatio
	return app
}

// WithProjectionNear adds the projection near to the builder
func (app *builder) WithProjectionNear(near float32) Builder {
	app.near = &near
	return app
}

// WithProjectionFar adds the projection far to the builder
func (app *builder) WithProjectionFar(far float32) Builder {
	app.far = &far
	return app
}

// WithIndex adds the index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Camera instance
func (app *builder) Now() (Camera, error) {
	if app.lookAtVariable == "" {
		return nil, errors.New("the lookAt variable is mandatory in order to build a Camera instance")
	}

	if app.eye == nil {
		return nil, errors.New("the lookAt eye is mandatory in order to build a Camera instance")
	}

	if app.center == nil {
		return nil, errors.New("the lookAt center is mandatory in order to build a Camera instance")
	}

	if app.up == nil {
		return nil, errors.New("the lookAt up is mandatory in order to build a Camera instance")
	}

	if app.projVariable == "" {
		return nil, errors.New("the projection variable is mandatory in order to build a Camera instance")
	}

	if app.fov == nil {
		return nil, errors.New("the projection fov is mandatory in order to build a Camera instance")
	}

	if app.aspectRatio == nil {
		return nil, errors.New("the projection aspectRatio is mandatory in order to build a Camera instance")
	}

	if app.near == nil {
		return nil, errors.New("the projection near is mandatory in order to build a Camera instance")
	}

	if app.far == nil {
		return nil, errors.New("the projection far is mandatory in order to build a Camera instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.lookAtVariable),
		[]byte(app.eye.String()),
		[]byte(app.center.String()),
		[]byte(app.up.String()),
		[]byte(app.projVariable),
		[]byte(strconv.FormatFloat(float64(*app.fov), 'f', 10, 32)),
		[]byte(strconv.FormatFloat(float64(*app.aspectRatio), 'f', 10, 32)),
		[]byte(strconv.FormatFloat(float64(*app.near), 'f', 10, 32)),
		[]byte(strconv.FormatFloat(float64(*app.far), 'f', 10, 32)),
		[]byte(strconv.Itoa(int(app.index))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hsh).CreatedOn(app.createdOn).Now()
	if err != nil {
		return nil, err
	}

	lookAt := createLookAt(app.lookAtVariable, *app.eye, *app.center, *app.up)
	projection := createProjection(app.projVariable, *app.fov, *app.aspectRatio, *app.near, *app.far)
	return createCamera(immutable, app.index, projection, lookAt), nil
}
