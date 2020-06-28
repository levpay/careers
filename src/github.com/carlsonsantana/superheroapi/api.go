package superheroapi

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		router,
	))
}
