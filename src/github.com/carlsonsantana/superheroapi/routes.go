package superheroapi

import "net/http"

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func GetRoutes() []Route {
	routes := []Route{
		Route{
			"POST",
			"/super",
			AddSuperHandler,
		},
		Route{
			"GET",
			"/super",
			ListSuperHandler,
		},
	}
	return routes
}
