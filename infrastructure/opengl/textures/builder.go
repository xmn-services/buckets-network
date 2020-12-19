package textures

import (
	"errors"
	"fmt"
	"image"
	image_color "image/color"
	"image/draw"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/xmn-services/buckets-network/domain/memory/worlds/scenes/nodes/models/materials/layers/layer/textures"
)

type builder struct {
	tex textures.Texture
}

func createBuilder() Builder {
	out := builder{
		tex: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTexture adds a texture to the builder
func (app *builder) WithTexture(tex textures.Texture) Builder {
	app.tex = tex
	return app
}

// Now builds a new Texture instance
func (app *builder) Now() (Texture, error) {
	if app.tex == nil {
		return nil, errors.New("the texture is mandatory in order to build a Texture instance")
	}

	pixels := app.tex.Pixels()
	uWidth, uHeight := pixels.Dimension()

	width := int(uWidth)
	height := int(uHeight)
	srcRGBA := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: width,
			Y: height,
		},
	})

	rows := pixels.All()
	for x, oneRow := range rows {
		elements := oneRow.All()
		for y, oneElemnt := range elements {
			color := oneElemnt.Color()
			alpha := oneElemnt.Alpha()
			rgbaColor := image_color.RGBA{
				R: color.Red(),
				G: color.Green(),
				B: color.Blue(),
				A: alpha,
			}

			srcRGBA.Set(x, y, rgbaColor)
		}
	}

	dimension := app.tex.Dimension()
	pos := dimension.Position()
	dim := dimension.Dimension()

	dstRGBA := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{
				X: pos.X(),
				Y: pos.Y(),
			},
			Max: image.Point{
				X: dim.X(),
				Y: dim.Y(),
			},
		},
	)

	if dstRGBA.Stride != dstRGBA.Rect.Size().X*4 {
		return nil, fmt.Errorf("the destination RGBA image has an unsupported stride")
	}

	draw.Draw(dstRGBA, dstRGBA.Bounds(), srcRGBA, image.Point{0, 0}, draw.Src)

	var identifier uint32
	gl.GenTextures(1, &identifier)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, identifier)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(dstRGBA.Rect.Size().X),
		int32(dstRGBA.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(dstRGBA.Pix),
	)

	return createTexture(app.tex, identifier), nil
}
