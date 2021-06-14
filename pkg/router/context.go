package router

import "github.com/bwmarrin/discordgo"

type Context struct {
	*discordgo.Session
	*discordgo.Message
	Vars map[interface{}]interface{}
}
