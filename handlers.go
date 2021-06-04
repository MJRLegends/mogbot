package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	mogbot "github.com/ChrisMcDearman/mogbot/mogbot"

	"github.com/bwmarrin/discordgo"
)

// Handlers is a a slice of all my handlers
var Handlers = []mogbot.Handler{
	func(b *mogbot.Bot) interface{} {
		return func(s *discordgo.Session, r *discordgo.Ready) {
			log.Printf("Logged in as %v\n", r.User)
			fmt.Println("-------------------------------------------------")
			if err := b.UpdateGameStatus(0, "DM to Contact Staff"); err != nil {
				log.Printf("Error updating status: %s", err)
				return
			}
		}
	},
	func(b *mogbot.Bot) interface{} {
		return func(s *discordgo.Session, g *discordgo.GuildCreate) {
			for _, r := range g.Roles {
				if err := b.DB.AddRole(*bot.NewRole(r)); err != nil {
					continue
				}
			}
			roles, err := b.DB.GetAllRoles()
			if err != nil {
				panic(err)
			}
			for _, m := range g.Members {
				nm := &mogbot.Member{m.User.ID, m.GuildID, []mogbot.Role{}}
				for _, r := range m.Roles {
					if fr, found := mogbot.FindRole(r, roles); found {
						nm.Roles = append(nm.Roles, *fr)
					}
				}
				if err := b.DB.AddMember(nm); err != nil {
					if err = b.DB.UpdateMember(nm); err != nil {
						continue
					}
				}
			}
			b.AddCommands(g.ID, Commands...)
			log.Printf("Added all members and commands for guild '%s'", g.ID)
		}
	},
	func(b *mogbot.Bot) interface{} {
		return func(s *discordgo.Session, m *discordgo.MessageCreate) {
			dmChannel, err := s.Channel(m.ChannelID)
			if err != nil {
				log.Printf("Error getting dm channel: %s", err)
				return
			}
			if dmChannel.Type != discordgo.ChannelTypeDM || m.Author.ID == s.State.User.ID {
				return
			}
			chName := strings.NewReplacer("#", " ").Replace(m.Author.String())
			guild, err := s.Guild(bot.GuildID)
			if err != nil {
				log.Printf("Error finding guild: %s", err)
				return
			}
			logFooter := &discordgo.MessageEmbedFooter{
				Text:    m.Author.String() + " | " + m.Author.ID,
				IconURL: guild.IconURL(),
			}
			logNewTicket := &discordgo.MessageEmbed{
				Title:     "New Ticket",
				Color:     0x2ecc71,
				Footer:    logFooter,
				Timestamp: time.Now().Format(time.RFC3339),
			}
			author := &discordgo.MessageEmbedAuthor{
				Name:    m.Author.String(),
				IconURL: m.Author.AvatarURL(""),
			}
			replyFooter := &discordgo.MessageEmbedFooter{
				Text:    "ID: " + m.Author.ID,
				IconURL: s.State.User.AvatarURL(""),
			}
			reply := &discordgo.MessageEmbed{
				Title:       "Message Sent",
				Description: m.Content,
				Color:       0xff4949,
				Author:      author,
				Footer:      replyFooter,
				Timestamp:   time.Now().Format(time.RFC3339),
			}
			mmFooter := &discordgo.MessageEmbedFooter{
				Text:    guild.Name + " | " + guild.ID,
				IconURL: guild.IconURL(),
			}
			mm := &discordgo.MessageEmbed{
				Title:       "Message Received",
				Description: m.Content,
				Color:       0x2ecc71,
				Author:      author,
				Footer:      mmFooter,
				Timestamp:   time.Now().Format(time.RFC3339),
			}
			s.ChannelMessageSendEmbed(bot.ModMailLogID, logNewTicket)
			s.ChannelMessageSendEmbed(m.ChannelID, reply)
			data := discordgo.GuildChannelCreateData{
				Name:     chName,
				ParentID: bot.ModMailCatID,
			}
			ch, err := s.GuildChannelCreateComplex(bot.GuildID, data)
			if err != nil {
				log.Print(err)
				return
			}
			_, err = s.ChannelMessageSendEmbed(ch.ID, mm)
			if err != nil {
				log.Print(err)
				return
			}
		}
	},
	//func(s *discordgo.Session, m *discordgo.MessageCreate) {
	//	_ = `^\[(?P<book>(?:\d\s*)?[A-Z]?[a-z]+)\s*(?P<chapter>\d+):(?P<verses>(?P<start>\d+)(?:-(?P<end>\d+))?)(?:\s(?P<version>[A-Z]?[a-z]+))?\]$`
	//},
}
