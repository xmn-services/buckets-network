package lists

type update struct {
	name        string
	description string
}

func createUpdateWithName(
	name string,
) Update {
	return createUpdateInternally(name, "")
}

func createUpdateWithDescription(
	description string,
) Update {
	return createUpdateInternally("", description)
}

func createUpdateWithNameAndDescription(
	name string,
	description string,
) Update {
	return createUpdateInternally(name, description)
}

func createUpdateInternally(
	name string,
	description string,
) Update {
	out := update{
		name:        name,
		description: description,
	}

	return &out
}

// HasName returns true if there is a name, false otherwise
func (obj *update) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *update) Name() string {
	return obj.name
}

// HasDescription returns true if there is a description, false otherwise
func (obj *update) HasDescription() bool {
	return obj.description != ""
}

// Description returns the description, if any
func (obj *update) Description() string {
	return obj.description
}
