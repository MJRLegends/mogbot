package handlers

import (
	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"github.com/bwmarrin/discordgo"
)

func OnModMailMessage(b *mogbot.Bot) interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//dmChannel, err := s.Channel(m.ChannelID)
		//if err != nil {
		//	log.Printf("Error getting dm channel: %s", err)
		//	return
		//}
		//if dmChannel.Type != discordgo.ChannelTypeDM || m.Author.ID == s.State.User.ID {
		//	return
		//}
		//chName := strings.NewReplacer("#", " ").Replace(m.Author.String())
		//guild, err := s.Guild(m.GuildID)
		//if err != nil {
		//	log.Printf("Error finding guild: %s", err)
		//	return
		//}
		//logFooter := &discordgo.MessageEmbedFooter{
		//	Text:    m.Author.String() + " | " + m.Author.ID,
		//	IconURL: guild.IconURL(),
		//}
		//logNewTicket := &discordgo.MessageEmbed{
		//	Title:     "New Ticket",
		//	Color:     0x2ecc71,
		//	Footer:    logFooter,
		//	Timestamp: time.Now().Format(time.RFC3339),
		//}
		//author := &discordgo.MessageEmbedAuthor{
		//	Name:    m.Author.String(),
		//	IconURL: m.Author.AvatarURL(""),
		//}
		//replyFooter := &discordgo.MessageEmbedFooter{
		//	Text:    "ID: " + m.Author.ID,
		//	IconURL: s.State.User.AvatarURL(""),
		//}
		//reply := &discordgo.MessageEmbed{
		//	Title:       "Message Sent",
		//	Description: m.Content,
		//	Color:       0xff4949,
		//	Author:      author,
		//	Footer:      replyFooter,
		//	Timestamp:   time.Now().Format(time.RFC3339),
		//}
		//mmFooter := &discordgo.MessageEmbedFooter{
		//	Text:    guild.Name + " | " + guild.ID,
		//	IconURL: guild.IconURL(),
		//}
		//mm := &discordgo.MessageEmbed{
		//	Title:       "Message Received",
		//	Description: m.Content,
		//	Color:       0x2ecc71,
		//	Author:      author,
		//	Footer:      mmFooter,
		//	Timestamp:   time.Now().Format(time.RFC3339),
		//}
		//s.ChannelMessageSendEmbed(bot.ModMailLogID, logNewTicket)
		//s.ChannelMessageSendEmbed(m.ChannelID, reply)
		//data := discordgo.GuildChannelCreateData{
		//	Name:     chName,
		//	ParentID: bot.ModMailCatID,
		//}
		//ch, err := s.GuildChannelCreateComplex(bot.GuildID, data)
		//if err != nil {
		//	log.Print(err)
		//	return
		//}
		//_, err = s.ChannelMessageSendEmbed(ch.ID, mm)
		//if err != nil {
		//	log.Print(err)
		//	return
		//}
	}
}
