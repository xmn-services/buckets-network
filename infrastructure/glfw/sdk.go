package glfw

import (
	"github.com/xmn-services/buckets-network/application/windows"
)

// NewBuilder creates a new glfw application builder
func NewBuilder() windows.Builder {
	return createBuilder()
}
