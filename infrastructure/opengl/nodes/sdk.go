package nodes

import (
	"github.com/go-gl/mathgl/mgl32"
	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	cameraBuilder := cameras.NewBuilder()
	programBuilder := programs.NewBuilder()
	modelBuilder := models.NewBuilder()
	return createBuilder(programBuilder, cameraBuilder, modelBuilder)
}

// Builder represents a node builder
type Builder interface {
	Create() Builder
	WithNode(node domain_nodes.Node) Builder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	Original() domain_nodes.Node
	Program() programs.Program
	Position() mgl32.Vec3
	Orientation() mgl32.Vec4
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
	Render(cameraIndex uint) error
}

// Content represents the node content
type Content interface {
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}
