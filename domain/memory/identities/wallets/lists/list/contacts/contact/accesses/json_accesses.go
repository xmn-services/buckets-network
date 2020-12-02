package accesses

// JSONAccesses represents a JSON access
type JSONAccesses struct {
	List []string `json:"access"`
}

func createJSONAccessesFromAccesses(ins Accesses) *JSONAccesses {
	lst := []string{}
	accesses := ins.All()
	for _, oneHash := range accesses {
		lst = append(lst, oneHash.String())
	}

	return createJSONAccesses(
		lst,
	)
}

func createJSONAccesses(
	lst []string,
) *JSONAccesses {
	out := JSONAccesses{
		List: lst,
	}

	return &out
}
