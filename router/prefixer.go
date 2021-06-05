package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Prefixer func(*discordgo.MessageCreate) (prefix string, ok bool)

func NewDefaultPrefixer(prefixes ...string) Prefixer {
	return func(m *discordgo.MessageCreate) (prefix string, ok bool) {
		for _, p := range prefixes {
			if strings.HasPrefix(m.Content, p) {
				return p, true
			}
		}
		return "", false
	}
}

func NewGuildPrefixer(prefixes map[string][]string) Prefixer {
	return func(m *discordgo.MessageCreate) (prefix string, ok bool) {
		if gp, ok := prefixes[m.GuildID]; ok {
			for _, p := range gp {
				if strings.HasPrefix(m.Content, p) {
					return p, true
				}
			}
			return "", false
		}
		return "", false
	}
}
