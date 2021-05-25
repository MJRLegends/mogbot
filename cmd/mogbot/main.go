package main

import (
	"flag"

	mogbot "github.com/captainmog/mogbot/internal/mogbot"

	"github.com/bwmarrin/discordgo"
	"github.com/captainmog/mogbot/internal/cache"
	"github.com/captainmog/mogbot/internal/exts"
	"github.com/captainmog/mogbot/internal/gorm"

	"github.com/captainmog/mogbot/pkg/memory"

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
	b := mogbot.NewBot(token, cache.NewCache(256, db))
	b.Identify.Intents = discordgo.IntentsAll
	for _, h := range exts.Handlers {
		b.AddHandler(h(b))
	}
	b.CommandHandlers = exts.CommandHandlers
	v := make(map[string]interface{})
	v["db"] = b.DB
	b.Start(b.Session, v)
	if err := b.Open(); err != nil {
		panic(err)
	}
	memory.Ticker()
	b.Wait()
}
