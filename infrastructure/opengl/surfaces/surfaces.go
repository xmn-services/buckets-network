package surfaces

import "github.com/xmn-services/buckets-network/infrastructure/opengl/surfaces/surface"

type surfaces struct {
	list []surface.Surface
}

func createSurfaces(
	list []surface.Surface,
) Surfaces {
	out := surfaces{
		list: list,
	}

	return &out
}

// All returns all surfaces
func (obj *surfaces) All() []surface.Surface {
	return obj.list
}
