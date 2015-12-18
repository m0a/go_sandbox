package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/m0a/go_sandbox/rest2/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(rootPath string) *mux.Router {

	var routes = Routes{
		Route{
			"files",
			"GET",
			"/api/files",
			handler.FileList(rootPath),
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
