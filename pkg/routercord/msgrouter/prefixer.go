package msgrouter

import "strings"

type Prefixer func(ctx *Context) (prefix string, ok bool)

func NewDefaultPrefixer(prefixes ...string) Prefixer {
	return func(ctx *Context) (prefix string, ok bool) {
		for _, p := range prefixes {
			if strings.HasPrefix(ctx.Content, p) {
				return p, true
			}
		}
		return "", false
	}
}

func NewGuildPrefixer(prefixes map[string][]string) Prefixer {
	return func(ctx *Context) (prefix string, ok bool) {
		if gp, ok := prefixes[ctx.GuildID]; ok {
			for _, p := range gp {
				if strings.HasPrefix(ctx.Content, p) {
					return p, true
				}
			}
			return "", false
		}
		return "", false
	}
}