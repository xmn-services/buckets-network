package contents

import "github.com/xmn-services/buckets-network/domain/memory/file/contents/content"

type builder struct {
	contentBuilder content.Builder
	contents       [][]byte
}

func createBuilder(
	contentBuilder content.Builder,
) Builder {
	out := builder{
		contentBuilder: contentBuilder,
		contents:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.contentBuilder)
}

// WithContents add contents to the builder
func (app *builder) WithContents(contents [][]byte) Builder {
	app.contents = contents
	return app
}

// Now builds a new Contents instance
func (app *builder) Now() (Contents, error) {
	lst := []content.Content{}
	mp := map[string]content.Content{}
	if app.contents != nil {
		for _, oneContent := range app.contents {
			content, err := app.contentBuilder.Create().WithContent(oneContent).Now()
			if err != nil {
				return nil, err
			}

			lst = append(lst, content)
			mp[content.Hash().String()] = content
		}
	}

	return createContents(lst, mp), nil
}
