package router

import (
	"github.com/bwmarrin/discordgo"
)

type Route struct {
	Name        string
	Aliases     []string
	Description string
	Handler     RouteHandler
	Parent      *Route
	Subroutes   []*Route
}

type RouteHandler func(*Context, []string)

func NewMessageRouter(prefix string) *Route {
	return &Route{Name: prefix}
}

func (r *Route) On(name string, handler RouteHandler) *Route {
	var nr = &Route{Name: name, Handler: handler, Parent: r}
	r.Subroutes = append(r.Subroutes, nr)
	return nr
}

func (r *Route) Alias(a ...string) *Route {
	r.Aliases = append(r.Aliases, a...)
	return r
}

func (r *Route) Execute(s *discordgo.Session, m *discordgo.Message, v map[string]interface{}) {

}
