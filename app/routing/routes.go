package routing

import (
    "net/http"
    handler "github.com/Indonesian-Numeral-Spellers/app/handlers"
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
        handler.SpellGet,
	},
	Route{
		"SpellPost",
        "POST",
        "/read",
        handler.SpellPost,
	},
}