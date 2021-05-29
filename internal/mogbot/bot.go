package mogbot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

type Bot struct {
	*discordgo.Session
	DB Database
	Plugins []Plugin
}

type Handler func(*Bot) interface{}

func NewBot(token string, db Database) *Bot {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	return &Bot{Session: s, DB: db}
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