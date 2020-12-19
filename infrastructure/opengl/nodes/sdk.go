package nodes

import (
	"time"

	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/renders"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	nodeBuilder := NewNodeBuilder()
	return createBuilder(nodeBuilder)
}

// NewNodeBuilder creates a new node builder instance
func NewNodeBuilder() NodeBuilder {
	spaceBuilder := spaces.NewBuilder()
	programBuilder := programs.NewBuilder()
	modelBuilder := models.NewBuilder()
	return createNodeBuilder(spaceBuilder, programBuilder, modelBuilder)
}

// Builder represents the nodes builder
type Builder interface {
	Create() Builder
	WithNodes(nodes []domain_nodes.Node) Builder
	Now() (Nodes, error)
}

// Nodes represents the nodes
type Nodes interface {
	All() []Node
	Camera(index uint) (cameras.Camera, spaces.Space, error)
	Render(delta time.Duration, camera cameras.Camera, globalSpace spaces.Space, renderApp renders.Application) error
}

// NodeBuilder represents a node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithNode(node domain_nodes.Node) NodeBuilder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	Original() domain_nodes.Node
	Space() spaces.Space
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
	Render(delta time.Duration, camera cameras.Camera, globalSpace spaces.Space, renderApp renders.Application) error
}

// Content represents the node content
type Content interface {
	IsModel() bool
	Model() models.Model
	IsCamera() bool
	Camera() cameras.Camera
}
