package main

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = Routes{
	Route{
		"StorageGet",
		"GET",
		"/api/storage/{key}",
		StorageGet,
	},
	Route{
		"StoragePut",
		"PUT",
		"/api/storage/{key}",
		StoragePut,
	},
	Route{
		"StoragePost",
		"POST",
		"/api/storage/{key}",
		StoragePost,
	},
	Route{
		"StorageDelete",
		"DELETE",
		"/api/storage/{key}",
		StorageDelete,
	},
	Route{
		"StorageOptions",
		"OPTIONS",
		"/api/storage/{key}",
		StorageOptions,
	},
}
