package commands

import "github.com/ChrisMcDearman/mogbot/router"

var Soundboard = &router.Route{
	Name:        "soundboard",
	Aliases:     []string{"sb"},
	Description: "Interacts with the soundboard",
	Handler:     nil,
	Parent:      nil,
	Subroutes:   nil,
	Middlewares: nil,
}
