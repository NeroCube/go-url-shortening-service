package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"URLIndex",
		"GET",
		"/urls",
		URLIndex,
	},
	Route{
		"URLCreate",
		"POST",
		"/urls",
		URLCreate,
	},
	Route{
		"URLShow",
		"GET",
		"/urls/{urlID}",
		URLShow,
	},
	Route{
		"URLRedirect",
		"GET",
		"/{tinyURL}",
		URLRedirect,
	},
}
