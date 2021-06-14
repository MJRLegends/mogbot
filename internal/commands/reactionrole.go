package commands

import "github.com/ChrisMcDearman/mogbot/pkg/router"

func ReactionRole() *router.Route {
	return &router.Route{
		Name:        "reactionrole",
		Aliases:     []string{"rr"},
		Description: "Interacts with reaction roles",
		Handler: func(ctx *router.Context, args []string) error {
			ctx.ChannelMessageSend(ctx.ChannelID, "You need to use this command with arguments.")
			return nil
		},
	}
}
