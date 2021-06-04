package commands

import (
	"github.com/ChrisMcDearman/mogbot/router"
)

var Echo = &router.Route{
	Name:        "echo",
	Description: "Echos a message",
	Handler: func(ctx *router.Context, args []string) {
		ctx.ChannelMessageSend(ctx.Message.ChannelID, args[0])
	},
	Middlewares: []router.Middleware{router.IsAdmin},
}
