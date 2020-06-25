package superheroapi

type SuperGroup struct {
	uuid string
	name string
}

type Super struct {
	uuid         string
	name         string
	fullName     string
	intelligence int
	power        int
	occupation   string
	image        string
	groups       []SuperGroup
	category     string
	relatives    []string
}
