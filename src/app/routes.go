package main

import (
	"net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "SpellGet",
        "GET",
        "/spell",
        SpellGet,
	},
	Route{
		"SpellPost",
        "POST",
        "/read",
        SpellPost,
	},
}