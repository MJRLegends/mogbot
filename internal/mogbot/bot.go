package mogbot

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"

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

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
