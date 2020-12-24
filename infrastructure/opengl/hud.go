package opengl

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type hud struct {
	id    *uuid.UUID
	nodes []HudNode
	mat   Material
}

func createHudWithNodes(
	id *uuid.UUID,
	nodes []HudNode,
) Hud {
	return createHudInternally(id, nodes, nil)
}

func createHudWithMaterial(
	id *uuid.UUID,
	mat Material,
) Hud {
	return createHudInternally(id, nil, mat)
}

func createHudWithNodesAndMaterial(
	id *uuid.UUID,
	nodes []HudNode,
	mat Material,
) Hud {
	return createHudInternally(id, nodes, mat)
}

func createHudInternally(
	id *uuid.UUID,
	nodes []HudNode,
	mat Material,
) Hud {
	out := hud{
		id:    id,
		nodes: nodes,
		mat:   mat,
	}

	return &out
}

// ID returns the id
func (obj *hud) ID() *uuid.UUID {
	return obj.id
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *hud) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes, if any
func (obj *hud) Nodes() []HudNode {
	return obj.nodes
}

// HasMaterial returns true if there is material, false otherwise
func (obj *hud) HasMaterial() bool {
	return obj.mat != nil
}

// Material returns the material, if any
func (obj *hud) Material() Material {
	return obj.mat
}

// Render renders the head-up display
func (obj *hud) Render(
	delta time.Duration,
	activeCamera WorldCamera,
	activeScene Scene,
) error {
	return nil
}
