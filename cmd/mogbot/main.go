package main

import (
	"flag"
	"fmt"
	cache2 "github.com/ChrisMcDearman/mogbot/internal/cache"
	commands2 "github.com/ChrisMcDearman/mogbot/internal/commands"
	gorm2 "github.com/ChrisMcDearman/mogbot/internal/gorm"
	handlers2 "github.com/ChrisMcDearman/mogbot/internal/handlers"

	"github.com/ChrisMcDearman/mogbot/router"

	"github.com/ChrisMcDearman/mogbot/mogbot"

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
	db, err := gorm2.NewConnection(*f)
	if err != nil {
		panic(err)
	}
	b := mogbot.New(token, cache2.New(256, db))
	b.AddRoutes(
		commands2.Ping(),
		commands2.Echo(),
	)
	b.Identify.Intents = discordgo.IntentsAll
	for _, h := range handlers2.Handlers {
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
