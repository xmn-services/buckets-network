package nodes

import (
	"github.com/xmn-services/buckets-network/domain/memory/worlds/math"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models"
	"github.com/xmn-services/buckets-network/libs/entities"
)

// Node represents a node
type Node interface {
	entities.Immutable
	Position() math.Vec3
	Right() math.Vec3
	Up() math.Vec3
	Content() Content
	HasChildren() bool
	Children() []Node
}

// Content represents the node content
type Content interface {
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}
