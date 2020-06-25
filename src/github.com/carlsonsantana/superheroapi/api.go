package superheroapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitPaths(router *mux.Router, routes []Route) {
	for _, route := range routes {
		router.HandleFunc(route.path, route.handler).Methods(route.method)
	}
}

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	InitPaths(router, GetRoutes())

	log.Fatal(http.ListenAndServe(":8080", router))
}
