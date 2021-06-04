package main

import (
	"flag"

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
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Bot token was not given.")
	}
	f := flag.String("dbmode", "DEV", "-dbmode=DEV")
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
		commands.Ping,
		commands.Echo,
	)
	b.Identify.Intents = discordgo.IntentsAll
	for _, h := range Handlers {
		b.AddHandler(h(b))
	}
	v := make(map[string]interface{})
	v["db"] = b.DB
	b.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		b.Execute(b.Session, router.NewDefaultPrefixer("!"), m.Message, v)
	})
	if err := b.Open(); err != nil {
		panic(err)
	}
	b.Wait()
}
