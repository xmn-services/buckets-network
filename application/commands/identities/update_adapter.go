package identities

import "net/url"

type updateAdapter struct {
	updateBuilder UpdateBuilder
}

func createUpdateAdapter(
	updateBuilder UpdateBuilder,
) UpdateAdapter {
	out := updateAdapter{
		updateBuilder: updateBuilder,
	}

	return &out
}

// URLValuesToUpdate converts url values to an Update instance
func (app *updateAdapter) URLValuesToUpdate(values url.Values) (Update, error) {
	builder := app.updateBuilder.Create()
	seed := values.Get("seed")
	if seed != "" {
		builder.WithSeed(seed)
	}

	name := values.Get("name")
	if name != "" {
		builder.WithName(name)
	}

	password := values.Get("password")
	if password != "" {
		builder.WithPassword(password)
	}

	root := values.Get("root")
	if root != "" {
		builder.WithRoot(root)
	}

	return builder.Now()
}

// UpdateToURLValues converts an Update instance to url values
func (app *updateAdapter) UpdateToURLValues(update Update) url.Values {
	values := url.Values{}
	if update.HasName() {
		values.Add("name", update.Name())
	}

	if update.HasRoot() {
		values.Add("root", update.Root())
	}

	if update.HasSeed() {
		values.Add("seed", update.Seed())
	}

	if update.HasPassword() {
		values.Add("password", update.Password())
	}

	return values
}
