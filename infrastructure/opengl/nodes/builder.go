package nodes

import (
	"errors"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	domain_nodes "github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/cameras"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/models"
	"github.com/xmn-services/buckets-network/infrastructure/opengl/programs"
)

type builder struct {
	programBuilder programs.Builder
	cameraBuilder  cameras.Builder
	modelBuilder   models.Builder
	node           domain_nodes.Node
}

func createBuilder(
	programBuilder programs.Builder,
	cameraBuilder cameras.Builder,
	modelBuilder models.Builder,
) Builder {
	out := builder{
		programBuilder: programBuilder,
		cameraBuilder:  cameraBuilder,
		modelBuilder:   modelBuilder,
		node:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.programBuilder,
		app.cameraBuilder,
		app.modelBuilder,
	)
}

// WithNode adds a node to the builder
func (app *builder) WithNode(node domain_nodes.Node) Builder {
	app.node = node
	return app
}

// Now builds a new Node instance
func (app *builder) Now() (Node, error) {
	if app.node == nil {
		return nil, errors.New("the node is mandatory in order to build a Node instance")
	}

	domainShaders := app.node.Shaders()
	prog, err := app.programBuilder.Create().WithShaders(domainShaders).Now()
	if err != nil {
		return nil, err
	}

	var content Content
	if app.node.HasContent() {
		nodeContent := app.node.Content()
		if nodeContent.IsCamera() {
			domainCamera := nodeContent.Camera()
			camera, err := app.cameraBuilder.Create().WithProgram(prog).WithCamera(domainCamera).Now()
			if err != nil {
				return nil, err
			}

			content = createContentWithCamera(camera)
		}

		if nodeContent.IsModel() {
			domainModel := nodeContent.Model()
			model, err := app.modelBuilder.Create().WithProgram(prog).WithModel(domainModel).Now()
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
			return createNodeWithContentAndChildren(app.node, prog, pos, orientation, content, nodeList), nil
		}

		return createNodeWithChildren(app.node, prog, pos, orientation, nodeList), nil
	}

	if content != nil {
		return createNodeWithContent(app.node, prog, pos, orientation, content), nil
	}

	return createNode(app.node, prog, pos, orientation), nil
}
