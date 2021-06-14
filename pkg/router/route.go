package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Route struct {
	Name        string
	Aliases     []string
	Description string
	Handler     RouteHandler
	Parent      *Route
	Subroutes   []*Route
	Middlewares []Middleware
}

type RouteHandler func(*Context, []string) error

func New() *Route {
	return &Route{}
}

func (r *Route) AddRoutes(subroutes ...*Route) *Route {
	for _, sr := range subroutes {
		sr.Parent = r
	}
	r.Subroutes = append(r.Subroutes, subroutes...)
	return r
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

func (r *Route) Desc(d string) *Route {
	r.Description = d
	return r
}

func (r *Route) AddMiddlewares(m ...Middleware) *Route {
	r.Middlewares = append(r.Middlewares, m...)
	return r
}

// Find finds a route with the given name
// It will return nil if nothing is found
//    name : name of route to find
func (r *Route) Find(name string) *Route {
	for _, sr := range r.Subroutes {
		if name == sr.Name {
			return sr
		}
		for _, a := range sr.Aliases {
			if name == a {
				return sr
			}
		}
	}
	return nil
}

// FindFull a full path of routes by searching through their subroutes
// Until the deepest match is found.
// It will return the route matched and the depth it was found at
//     args : path of route you wish to find
//            ex. FindFull(command, subroute1, subroute2, nonexistent)
//            will return the deepest found match, which will be subroute2
func (r *Route) FindFull(args []string) (*Route, int) {
	nr := r
	i := 0
	for _, v := range args {
		if rt := nr.Find(v); rt != nil {
			nr = rt
			i++
		} else {
			break
		}
	}
	return nr, i
}

func (r *Route) ApplyMiddlewares() RouteHandler {
	h := r.Handler
	for _, m := range r.Middlewares {
		h = m(h)
	}
	return h
}

func (r *Route) Execute(s *discordgo.Session, parser Parser, prefixer Prefixer, m *discordgo.MessageCreate, v map[interface{}]interface{}) error {
	p, ok := prefixer(m)
	if !ok {
		return nil
	}
	args, err := parser(strings.TrimPrefix(m.Content, p))
	if err != nil {
		return err
	}
	rt, d := r.FindFull(args)
	if d == 0 {
		return nil
	}
	return rt.ApplyMiddlewares()(&Context{Session: s, Message: m.Message, Vars: v}, args[d:])
}
