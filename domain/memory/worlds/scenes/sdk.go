package scenes

import (
	uuid "github.com/satori/go.uuid"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/huds"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the scene builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithIndex(index uint) Builder
	WithHud(hud huds.Hud) Builder
	WithNodes(nodes []nodes.Node) Builder
	Now() (Scene, error)
}

// Scene represents a scene
type Scene interface {
	ID() *uuid.UUID
	Index() uint
	Hud() huds.Hud
	Nodes() []nodes.Node
}
