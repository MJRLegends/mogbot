package commands

import (
	"github.com/ChrisMcDearman/mogbot/router"
)

func Echo() *router.Route {
	return &router.Route{
		Name:        "echo",
		Description: "Echos a message",
		Handler: func(ctx *router.Context, args []string) error {
			_, err := ctx.ChannelMessageSend(ctx.Message.ChannelID, args[0])
			return err
		},
		Middlewares: []router.Middleware{router.IsAdmin},
	}
}
