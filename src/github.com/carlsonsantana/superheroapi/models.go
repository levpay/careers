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

type SuperAPIResponsePowerStatus struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type SuperAPIResponseBiography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceBirth      string   `json:"place-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type SuperAPIResponseAppearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}

type SuperAPIResponseWork struct {
	Occupation string `json:"occupation"`
	Case       string `json:"case"`
}

type SuperAPIResponseConnections struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}

type SuperAPIResponseImage struct {
	URL string `json:"url"`
}

type SuperAPIResponseSuper struct {
	Id          string                      `json:"id"`
	Name        string                      `json:"name"`
	Powerstats  SuperAPIResponsePowerStatus `json:"powerstats"`
	Biography   SuperAPIResponseBiography   `json:"biography"`
	Appearance  SuperAPIResponseAppearance  `json:"appearance"`
	Work        SuperAPIResponseWork        `json:"work"`
	Connections SuperAPIResponseConnections `json:"connections"`
	Image       SuperAPIResponseImage       `json:"image"`
}

type SuperAPIResponse struct {
	Response string                  `json:"response"`
	Results  []SuperAPIResponseSuper `json:"results"`
	Error    string                  `json:"error"`
}
