package commands

import "github.com/ChrisMcDearman/mogbot/router"

func Soundboard() *router.Route {
	var soundboard = &router.Route{
		Name:        "soundboard",
		Aliases:     []string{"sb"},
		Description: "Interacts with the soundboard",
		Handler: func(ctx *router.Context, args []string) error {
			_, err := ctx.ChannelMessageSend(ctx.Message.ChannelID, "Invalid command invoke: You must use a subcommand with this command.")
			return err
		},
	}
	soundboard.AddRoutes(Play())
	return soundboard
}

func Play() *router.Route {
	return &router.Route{
		Name:        "play",
		Aliases:     []string{"p"},
		Description: "Plays a sound on the soundboard",
		Handler: func(ctx *router.Context, args []string) error {
			//sb := ctx.Vars["player"]
			return nil
		},
	}
}
