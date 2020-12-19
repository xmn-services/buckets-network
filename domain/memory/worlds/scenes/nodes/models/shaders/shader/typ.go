package shader

type typ struct {
	isVertex   bool
	isFragment bool
}

func createTypeWithVertex() Type {
	return createTypeInternally(true, false)
}

func createTypeWithFragment() Type {
	return createTypeInternally(false, true)
}

func createTypeInternally(
	isVertex bool,
	isFragment bool,
) Type {
	out := typ{
		isVertex:   isVertex,
		isFragment: isFragment,
	}

	return &out
}

// IsVertex returns true if the shader is a vertex shader, false otherwise
func (obj *typ) IsVertex() bool {
	return obj.isVertex
}

// IsFragment returns true if the shader is a fragment shader, false otherwise
func (obj *typ) IsFragment() bool {
	return obj.isFragment
}
