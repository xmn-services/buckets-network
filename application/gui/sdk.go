package gui

import (
	"github.com/xmn-services/buckets-network/domain/memory/windows"
	"github.com/xmn-services/buckets-network/domain/memory/worlds"
)

// Application represents a gui application
type Application interface {
	Execute(win windows.Window, world worlds.World) error
}
