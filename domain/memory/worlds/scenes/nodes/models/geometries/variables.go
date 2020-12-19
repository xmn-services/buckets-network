package geometries

type variables struct {
	vertex  string
	texture string
}

func createVariables(
	vertex string,
	texture string,
) Variables {
	out := variables{
		vertex:  vertex,
		texture: texture,
	}

	return &out
}

// VertexCoordinates returns the vertex coordinates
func (obj *variables) VertexCoordinates() string {
	return obj.vertex
}

// TextureCoordinates returns the texture coordinates
func (obj *variables) TextureCoordinates() string {
	return obj.texture
}
