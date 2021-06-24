package handlers

import (
	"fmt"
	"log"

	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"github.com/bwmarrin/discordgo"
)

func OnReady(b *mogbot.Bot) interface{} {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		m := fmt.Sprintf("Logged in as %v\n", r.User)
		log.Print(m)
		for i := -19; i < len(m); i++ {
			fmt.Print("-")
		}
		fmt.Println()
		//if err := b.UpdateGameStatus(0, "DM to Contact Staff"); err != nil {
		//	log.Printf("Error updating status: %s", err)
		//	return
		//}
	}
}
