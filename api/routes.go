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
		"deploy_create",
		"POST",
		"/deploy/stack/create",
		deploy_create,
	},
	Route{
		"deploy_remove",
		"POST",
		"/deploy/stack/remove",
		deploy_remove,
	},
	Route{
		"deploy_status",
		"POST",
		"/deploy/stack/{name}/status",
		deploy_status,
	},
}
