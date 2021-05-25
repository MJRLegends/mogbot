package mogbot

import (
	"github.com/captainmog/mogbot/pkg/routercord/slashrouter"
)

type Bot struct {
	*slashrouter.Router
	DB Database
}

type Handler func(*Bot) interface{}

func NewBot(token string, db Database) *Bot {
	return &Bot{slashrouter.NewRouter(token), db}
}
