package commands

import (
	"fmt"
	"time"

	"github.com/ChrisMcDearman/mogbot/pkg/router"
)

func Ping() *router.Route {
	return &router.Route{
		Name:        "ping",
		Description: "Pings the bot. Responds with pong and latency.",
		Handler: func(ctx *router.Context, args []string) error {
			_, err := ctx.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("Pong %dms!", ctx.HeartbeatLatency()/time.Millisecond))
			return err
		},
	}

}
