package router

type Middleware func(RouteHandler) RouteHandler

func IsAdmin(h RouteHandler) RouteHandler {
	return func(ctx *Context, args []string) error {
		if IsChannelAdmin(ctx.Session, ctx.Message.ChannelID, ctx.Author.ID) {
			return h(ctx, args)
		} else {
			_, err := ctx.ChannelMessageSend(ctx.Message.ChannelID, "You must be an admin to run this command!")
			return err
		}
	}
}
