package nodes

import (
	"errors"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/spaces"
)

type nodeBuilder struct {
	spaceBuilder   spaces.Builder
	programBuilder programs.Builder
	modelBuilder   models.Builder
	node           domain_nodes.Node
}

func createNodeBuilder(
	spaceBuilder spaces.Builder,
	programBuilder programs.Builder,
	modelBuilder models.Builder,
) NodeBuilder {
	out := nodeBuilder{
		spaceBuilder:   spaceBuilder,
		programBuilder: programBuilder,
		modelBuilder:   modelBuilder,
		node:           nil,
	}

	return &out
}

// Create initializes the node builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder(
		app.spaceBuilder,
		app.programBuilder,
		app.modelBuilder,
	)
}

// WithNode adds a node to the builder
func (app *nodeBuilder) WithNode(node domain_nodes.Node) NodeBuilder {
	app.node = node
	return app
}

// Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.node == nil {
		return nil, errors.New("the node is mandatory in order to build a Node instance")
	}

	var content Content
	if app.node.HasContent() {
		nodeContent := app.node.Content()
		if nodeContent.IsCamera() {
			cam := nodeContent.Camera()
			content = createContentWithCamera(cam)
		}

		if nodeContent.IsModel() {
			domainModel := nodeContent.Model()
			model, err := app.modelBuilder.Create().WithModel(domainModel).Now()
			if err != nil {
				return nil, err
			}

			content = createContentWithModel(model)
		}
	}

	domainPos := app.node.Position()
	domainOrientation := app.node.Orientation()
	domainDirection := domainOrientation.Direction()
	radAngle := float32(domainOrientation.Angle() * math.Pi / 180)

	pos := mgl32.Vec3{
		domainPos.X(),
		domainPos.Y(),
		domainPos.Z(),
	}

	orientation := mgl32.Vec4{
		domainDirection.X(),
		domainDirection.Y(),
		domainDirection.Z(),
		radAngle,
	}

	space, err := app.spaceBuilder.Create().WithPosition(pos).WithOrientation(orientation).Now()
	if err != nil {
		return nil, err
	}

	if app.node.HasChildren() {
		nodeList := []Node{}
		children := app.node.Children()
		for _, oneDomainNode := range children {
			node, err := app.Create().WithNode(oneDomainNode).Now()
			if err != nil {
				return nil, err
			}

			nodeList = append(nodeList, node)
		}

		if content != nil {
			return createNodeWithContentAndChildren(app.node, space, content, nodeList), nil
		}

		return createNodeWithChildren(app.node, space, nodeList), nil
	}

	if content != nil {
		return createNodeWithContent(app.node, space, content), nil
	}

	return createNode(app.node, space), nil
}
