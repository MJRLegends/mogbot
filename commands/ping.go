package commands

import (
	"fmt"
	"time"

	"github.com/ChrisMcDearman/mogbot/router"
)

var Ping = &router.Route{
	Name:        "ping",
	Description: "Pings the bot. Responds with pong and latency.",
	Handler: func(ctx *router.Context, args []string) {
		ctx.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("Pong %dms!", ctx.HeartbeatLatency()/time.Millisecond))
	},
}