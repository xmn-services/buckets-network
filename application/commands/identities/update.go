package identities

type update struct {
	seed     string
	name     string
	password string
	root     string
}

func createUpdate(
	seed string,
	name string,
	password string,
	root string,
) Update {
	out := update{
		seed:     seed,
		name:     name,
		password: password,
		root:     root,
	}

	return &out
}

// HasSeed returns true if there is a seed, false otherwise
func (obj *update) HasSeed() bool {
	return obj.seed != ""
}

// Seed returns the seed, if any
func (obj *update) Seed() string {
	return obj.seed
}

// HasName returns true if there is a name, false otherwise
func (obj *update) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *update) Name() string {
	return obj.name
}

// HasPassword returns true if there is a password, false otherwise
func (obj *update) HasPassword() bool {
	return obj.password != ""
}

// Password returns the password, if any
func (obj *update) Password() string {
	return obj.password
}

// HasRoot returns true if there is a root, false otherwise
func (obj *update) HasRoot() bool {
	return obj.root != ""
}

// Root returns the root, if any
func (obj *update) Root() string {
	return obj.root
}
