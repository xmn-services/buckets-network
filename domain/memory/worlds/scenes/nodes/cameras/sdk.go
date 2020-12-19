package cameras

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/math/fl32"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	immutableBuilder := entities.NewImmutableBuilder()
	return createBuilder(hashAdapter, immutableBuilder)
}

// Builder represents a camera builder
type Builder interface {
	Create() Builder
	WithLookAtVariable(lookAtVariable string) Builder
	WithLookAtEye(eye fl32.Vec3) Builder
	WithLookAtCenter(center fl32.Vec3) Builder
	WithLookAtUp(up fl32.Vec3) Builder
	WithProjectionVariable(projVariable string) Builder
	WithProjectionFieldofView(fov float32) Builder
	WithProjectionAspectRatio(aspectRatio float32) Builder
	WithProjectionNear(near float32) Builder
	WithProjectionFar(far float32) Builder
	WithIndex(index uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	entities.Immutable
	Index() uint
	Projection() Projection
	LookAt() LookAt
}

// LookAt represents the direction where the camera looks at
type LookAt interface {
	Variable() string
	Eye() fl32.Vec3
	Center() fl32.Vec3
	Up() fl32.Vec3
}

// Projection represents the camera projection
type Projection interface {
	Variable() string
	FieldOfView() float32
	AspectRation() float32
	Near() float32
	Far() float32
}
