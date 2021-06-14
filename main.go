package main

import (
	"flag"
	"fmt"

	"github.com/ChrisMcDearman/mogbot/handlers"

	"github.com/ChrisMcDearman/mogbot/router"

	"github.com/ChrisMcDearman/mogbot/commands"

	"github.com/ChrisMcDearman/mogbot/gorm"

	"github.com/ChrisMcDearman/mogbot/mogbot"

	"github.com/ChrisMcDearman/mogbot/cache"

	"github.com/bwmarrin/discordgo"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags)
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	token := os.Getenv("PROD")
	if token == "" {
		log.Fatal("Bot token was not given.")
	}
	f := flag.String("dbmode", "DEV_DB", "-dbmode=DEV_DB")
	flag.Parse()
	if *f == "" {
		log.Fatal("DB mode was not given.")
	}
	db, err := gorm.NewConnection(*f)
	if err != nil {
		panic(err)
	}
	b := mogbot.New(token, cache.New(256, db))
	b.AddRoutes(
		commands.Ping(),
		commands.Echo(),
	)
	b.Identify.Intents = discordgo.IntentsAll
	for _, h := range handlers.Handlers {
		b.AddHandler(h(b))
	}
	v := make(map[interface{}]interface{})
	v["db"] = b.DB
	b.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if err := b.Execute(b.Session, router.NewDefaultParser(), router.NewDefaultPrefixer("!"), m, v); err != nil {
			log.Printf("Command error: %s", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error invoking command: %s", err))
		}
	})
	if err := b.Open(); err != nil {
		panic(err)
	}
	b.Wait()
}
