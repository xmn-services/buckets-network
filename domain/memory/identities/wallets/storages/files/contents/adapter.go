package contents

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// ToJSON converts Content instance to json
func (app *adapter) ToJSON(content Content) *JSONContent {
	return createJSONContentFromContent(content)
}

// ToContent converts json to Content instance
func (app *adapter) ToContent(js *JSONContent) (Content, error) {
	return createContentFromJSON(js)
}
