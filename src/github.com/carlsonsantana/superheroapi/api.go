package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initPaths(router *mux.Router, routes []Route) {
	for _, route := range routes {
		router.HandleFunc(route.path, route.handler).Methods(route.method)
	}
}

func initAPI() {
	router := mux.NewRouter().StrictSlash(true)

	initPaths(router, getRoutes())

	log.Fatal(http.ListenAndServe(":8080", router))
}
