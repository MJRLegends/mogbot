package slashrouter

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Router struct {
	*discordgo.Session
	CommandHandlers map[string]CommandHandler
}

type Context struct {
	*discordgo.Session
	*discordgo.Interaction
	Vars map[string]interface{}
}

type CommandHandler func(*Context)

func NewRouter(token string) *Router {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	return &Router{Session: session, CommandHandlers: make(map[string]CommandHandler)}
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

func (r *Router) Wait() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Print("Shutting down gracefully.")
	r.Close()
}
