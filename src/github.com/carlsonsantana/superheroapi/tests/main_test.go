package tests

import (
	"os"
	"testing"

	"github.com/gorilla/mux"

	"github.com/carlsonsantana/superheroapi"
)

var TestRouter *mux.Router
var TestRoutes []superheroapi.Route

func setup() {
	TestRoutes = superheroapi.GetRoutes()
	TestRouter = mux.NewRouter().StrictSlash(true)
	superheroapi.InitPaths(TestRouter, TestRoutes)
}

func shutdown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
