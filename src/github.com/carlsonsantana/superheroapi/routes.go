package main

import "net/http"

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func getRoutes() []Route {
	routes := []Route{}
	return routes
}
