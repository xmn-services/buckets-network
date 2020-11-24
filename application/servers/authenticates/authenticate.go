package authenticates

type authenticate struct {
	name     string
	password string
	seed     string
}

func createAuthenticate(
	name string,
	password string,
	seed string,
) Authenticate {
	out := authenticate{
		name:     name,
		password: password,
		seed:     seed,
	}

	return &out
}

// Name returns the name
func (obj *authenticate) Name() string {
	return obj.name
}

// Password returns the password
func (obj *authenticate) Password() string {
	return obj.password
}

// Seed returns the seed
func (obj *authenticate) Seed() string {
	return obj.seed
}
