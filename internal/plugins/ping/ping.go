package ping

import (
	"github.com/captainmog/mogbot/internal/mogbot"
	"github.com/captainmog/mogbot/pkg/router"
)

type Ping struct {
	MessageRouter *router.Route
	SlashRouter   map[string]router.SlashCommandHandler
	Handlers      []mogbot.Handler
}

func (p *Ping) RegisterPlugin(b mogbot.Bot) {
	b.Plugins = append(b.Plugins, p)
}

func (p *Ping) PluginInfo() mogbot.PluginInfo {
	return mogbot.PluginInfo{Name: "Ping"}
}

func NewPing() *Ping {
	return &Ping{MessageRouter: router.NewMessageRouter("!")}
}