package mogbot

import (
	"log"
	"os"
	"os/signal"

	"github.com/ChrisMcDearman/mogbot/router"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	*discordgo.Session
	DB Database
	//router.SlashRouter
	*router.Route
	Handlers []Handler
}

type Handler func(*Bot) interface{}

func New(token string, db Database) *Bot {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	return &Bot{Session: s, DB: db, Route: router.New()}
}

func (b *Bot) Wait() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Print("Shutting down gracefully...")
	if err := b.Close(); err != nil {
		panic(err)
	}
}
