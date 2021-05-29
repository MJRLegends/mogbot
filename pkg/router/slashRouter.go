package router

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type SlashRouter map[string]SlashCommandHandler

type SlashCommandHandler func(*Context)

func NewSlashRouter() SlashRouter {
	return make(map[string]SlashCommandHandler)
}

type Context struct {
	*discordgo.Session
	*discordgo.Message
	*discordgo.Interaction
	Vars map[string]interface{}
}



func (r *Router) AddCommands(guildID string, cmds ...*discordgo.ApplicationCommand) {
	oldCmds, err := r.ApplicationCommands(r.State.User.ID, guildID)
	if err != nil {
		log.Printf("Error retrieving commands for Guild '%s': %s", guildID, err)
		return
	}
	for _, c := range cmds {
		if FindCommand(c.Name, oldCmds) {
			continue
		}
		_, err = r.ApplicationCommandCreate(r.State.User.ID, guildID, c)
		log.Printf("Added command '%s'", c.Name)
		if err != nil {
			log.Panicf("Cannot create '%s' command: %s", c.Name, err)
		}
	}
}

func (r *Router) Start(s *discordgo.Session, v map[string]interface{}) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Printf("InteractionCreate: '%s'", i.Data.Name)
		if h, ok := r.CommandHandlers[i.Data.Name]; ok {
			log.Printf("Application command invoke: '%s'", i.Data.Name)
			h(&Context{s, i.Interaction, v})
		}
	})
	log.Print("Added router handler to session")
}


