package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route ... < Some lines that describe your function>
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ... < Some lines that describe your function>
type Routes []Route

// NewRouter ... < Some lines that describe your function>
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	{
		"beers",
		"GET",
		"/beers",
		beers,
	},
	{
		"BeerDetail",
		"GET",
		"/beers/{beerID}",
		BeerDetail,
	},
	{
		"beers",
		"POST",
		"/beers",
		beersIng,
	},
	{
		"beersBoxprice",
		"get",
		"/beers/{beerID}/boxprice",
		beersBoxprice,
	},
}
