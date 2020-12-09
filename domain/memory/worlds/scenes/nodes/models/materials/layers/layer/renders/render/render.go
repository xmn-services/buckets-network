package render

import (
	"time"

	"github.com/xmn-services/buckets-network/domain/memory/worlds/shapes/rectangles"
	"github.com/xmn-services/buckets-network/libs/entities"
	"github.com/xmn-services/buckets-network/libs/hash"
)

type render struct {
	immutable entities.Immutable
	opacity   float64
	viewport  rectangles.Rectangle
	content   Content
}

func createRender(
	immutable entities.Immutable,
	opacity float64,
	viewport rectangles.Rectangle,
	content Content,
) Render {
	out := render{
		immutable: immutable,
		opacity:   opacity,
		viewport:  viewport,
		content:   content,
	}

	return &out
}

// Hash returns the hash
func (obj *render) Hash() hash.Hash {
	return obj.immutable.Hash()
}

// Opacity returns the opacity
func (obj *render) Opacity() float64 {
	return obj.opacity
}

// Viewport returns the viewport
func (obj *render) Viewport() rectangles.Rectangle {
	return obj.viewport
}

// CreatedOn returns the creation time
func (obj *render) CreatedOn() time.Time {
	return obj.immutable.CreatedOn()
}

// HasContent returns true if there is content, false otherwise
func (obj *render) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *render) Content() Content {
	return obj.content
}
