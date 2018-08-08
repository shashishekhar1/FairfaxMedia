package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route Object to store Routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array of routes
type Routes []Route

// NewRouter reads the Route array and cretaes the routes
func NewRouter() *mux.Router {
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

var routes = Routes{
	Route{
		"CreateArticle",
		"POST",
		"/articles",
		CreateArticle,
	},
	Route{
		"GetArticle",
		"GET",
		"/articles/{id}",
		GetArticle,
	},
	Route{
		"GetDateBasedTagName",
		"GET",
		"/tags/{tagName}/{date}",
		GetDateBasedTagName,
	},
}
