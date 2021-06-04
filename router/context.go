package router

import "github.com/bwmarrin/discordgo"

type Context struct {
	*discordgo.Session
	*discordgo.Message
	*discordgo.Interaction
	Vars map[interface{}]interface{}
}
