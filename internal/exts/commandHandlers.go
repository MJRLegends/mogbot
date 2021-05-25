package exts

import (
	"fmt"
	"github.com/captainmog/mogbot/pkg/routercord/slashrouter"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// CommandHandlers is a slice of all my command handlers
var CommandHandlers = map[string]slashrouter.CommandHandler{
	"ping": func(c *slashrouter.Context) {
		err := c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: fmt.Sprintf("Pong %dms!", c.HeartbeatLatency()/time.Millisecond),
			},
		})
		if err != nil {
			log.Printf("Error sending message: %s", err)
		}
	},
	"move": func(c *slashrouter.Context) {
		if !slashrouter.IsAdmin(c.Session, c.GuildID, c.Member) {
			c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: "You lack the permissions to run this command",
				},
			})
			return
		}
		g, err := c.State.Guild(c.GuildID)
		if err != nil {
			log.Printf("Error retrieving guild '%s': %s", c.GuildID, err)
			return
		}
		var from *discordgo.Channel
		if ok, o := slashrouter.FindOption("from", c.Data.Options); ok {
			from = o.ChannelValue(c.Session)
		}
		var to *discordgo.Channel
		if ok, o := slashrouter.FindOption("to", c.Data.Options); ok {
			to = o.ChannelValue(c.Session)
		}
		for _, vs := range g.VoiceStates {
			if vs.ChannelID == from.ID {
				if err := c.GuildMemberMove(c.GuildID, vs.UserID, &to.ID); err != nil {
					log.Printf("Error moving member: %s", err)
					return
				}
			}
		}
		c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	},
	//"warn": func(c *routercord.Context) {
	//	var u *discordgo.User
	//	if ok, o := routercord.FindOption("user", c.Data.Options); ok {
	//		u = o.UserValue(c.Session)
	//	}
	//	s := c.Vars["db"].(mogbot.MemberService)
	//	m, err := s.GetMember(u.ID)
	//	if err != nil {
	//		c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
	//			Type: discordgo.InteractionResponseChannelMessageWithSource,
	//			Data: &discordgo.InteractionApplicationCommandResponseData{
	//				Content: fmt.Sprintf("Error executing command: %s", err),
	//			},
	//		})
	//	}
	//	var reason string
	//	if ok, o := routercord.FindOption("reason", c.Data.Options); ok {
	//		reason = o.StringValue()
	//	}
	//	c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
	//		Type: discordgo.InteractionResponseChannelMessageWithSource,
	//		Data: &discordgo.InteractionApplicationCommandResponseData{
	//			Content: fmt.Sprintf("Warned %s for **'%s'**", u.Mention(), reason),
	//		},
	//	})
	//},
	"ban": func(c *slashrouter.Context) {
		if !slashrouter.CanBan(c.Session, c.GuildID, c.Member) {
			err := c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: fmt.Sprintf("You lack the permissions to run this command."),
				},
			})
			if err != nil {
				log.Printf("Error sending message: %s", err)
			}
			return
		}
		var user *discordgo.User
		var reason string = ""
		var days int = 0
		if found, o := slashrouter.FindOption("user", c.Data.Options); found {
			user = o.UserValue(c.Session)
		}
		if found, o := slashrouter.FindOption("reason", c.Data.Options); found {
			reason = o.StringValue()
		}
		if found, o := slashrouter.FindOption("days", c.Data.Options); found {
			days = int(o.IntValue())
		}
		if err := c.GuildBanCreateWithReason(c.GuildID, user.ID, reason, days); err != nil {
			log.Printf("Error banning user '%s': %s", c.Data.Options[0].UserValue(c.Session).ID, err)
		}
		if err := c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: fmt.Sprintf("%s was banned", user.Mention()),
			},
		}); err != nil {
			log.Printf("Error responding to interaction: %s", err)
		}
	},
	"embed": func(c *slashrouter.Context) {
		if !slashrouter.IsAdmin(c.Session, c.GuildID, c.Member) {
			c.ChannelMessageSend(c.ChannelID, "You lack the permissions to run this command.")
			return
		}
		switch c.Data.Options[0].Name {
		case "create":
			var channelID string
			var title string
			var desc string
			var color int
			var author discordgo.MessageEmbedAuthor
			var thumbnail discordgo.MessageEmbedThumbnail
			var footer discordgo.MessageEmbedFooter
			var timestamp string
			for _, o := range c.Data.Options[0].Options {
				switch o.Name {
				case "title":
					title = fmt.Sprintf("%v", o.Value)
				case "description":
					desc = fmt.Sprintf("%v", o.Value)
				case "color":
					color = int(o.IntValue())

				case "author":
					u := o.UserValue(c.Session)
					author.Name = u.String()
					author.IconURL = u.AvatarURL("")
				case "thumbnail":
					thumbnail.URL = o.StringValue()
				case "footer":
					for _, fo := range o.Options {
						switch fo.Name {
						case "text":
							footer.Text = fo.StringValue()
						case "icon":
							if fo.StringValue() == "server" {
								g, err := c.Guild(c.GuildID)
								if err != nil {
									log.Printf("Error retrieving guild: %s", err)
									return
								}
								footer.IconURL = g.IconURL()
							} else {
								footer.IconURL = fo.StringValue()
							}
						}
					}
				case "timestamp":
					if o.BoolValue() {
						timestamp = time.Now().Format(time.RFC3339)
					}
				case "channel":
					channelID = o.ChannelValue(c.Session).ID
				}
			}
			embed := &discordgo.MessageEmbed{
				Title:       title,
				Description: desc,
				Color:       color,
				Author:      &author,
				Thumbnail:   &thumbnail,
				Footer:      &footer,
				Timestamp:   timestamp,
			}
			if channelID == "" {

			}
			if _, err := c.ChannelMessageSendEmbed(channelID, embed); err != nil {
				log.Printf("Error sending message: %s", err)
			}
			if err := c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			}); err != nil {
				log.Printf("Error responding to interaction: %s", err)
			}
		case "edit":
			found, o := slashrouter.FindOption("channel", c.Data.Options[0].Options)
			if !found {
				log.Printf("Unable to find channel option")
				return
			}
			channelID := o.ChannelValue(c.Session).ID
			found, o = slashrouter.FindOption("messageid", c.Data.Options[0].Options)
			if !found {
				log.Printf("Unable to find messageID option")
				return
			}
			messageID := o.StringValue()
			m, err := c.ChannelMessage(channelID, messageID)
			if err != nil {
				log.Printf("Error finding message: %s", err)
				return
			}
			title := m.Embeds[0].Title
			desc := m.Embeds[0].Description
			color := m.Embeds[0].Color
			author := m.Embeds[0].Author
			thumbnail := m.Embeds[0].Thumbnail
			footer := m.Embeds[0].Footer
			timestamp := m.Embeds[0].Timestamp
			if found, o := slashrouter.FindOption("title", c.Data.Options[0].Options); found {
				title = fmt.Sprintf("%v", o.Value)
			}
			if found, o := slashrouter.FindOption("description", c.Data.Options[0].Options); found {
				desc = fmt.Sprintf("%v", o.Value)
			}
			if found, o := slashrouter.FindOption("color", c.Data.Options[0].Options); found {
				color = int(o.IntValue())
			}
			if found, o := slashrouter.FindOption("author", c.Data.Options[0].Options); found {
				u := o.UserValue(c.Session)
				author.Name = u.String()
				author.IconURL = u.AvatarURL("")
			}
			if found, o := slashrouter.FindOption("thumbnail", c.Data.Options[0].Options); found {
				thumbnail.URL = o.StringValue()
			}
			if found, o := slashrouter.FindOption("footertext", c.Data.Options[0].Options); found {
				footer.Text = o.StringValue()
			}
			if found, o := slashrouter.FindOption("footericon", c.Data.Options[0].Options); found {
				if o.StringValue() == "server" {
					g, err := c.Guild(c.GuildID)
					if err != nil {
						log.Printf("Error retrieving guild: %s", err)
						return
					}
					footer.IconURL = g.IconURL()
				} else {
					footer.IconURL = o.StringValue()
				}
			}
			if found, o := slashrouter.FindOption("timestamp", c.Data.Options[0].Options); found {
				if o.BoolValue() {
					timestamp = time.Now().Format(time.RFC3339)
				}
			}
			embed := &discordgo.MessageEmbed{
				Title:       title,
				Description: desc,
				Color:       color,
				Author:      author,
				Thumbnail:   thumbnail,
				Footer:      footer,
				Timestamp:   timestamp,
			}
			if _, err := c.ChannelMessageEditEmbed(channelID, messageID, embed); err != nil {
				log.Printf("Error sending message: %s", err)
			}
			err = c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			})
			if err != nil {
				log.Printf("Interaction response error: %s", err)
			}
		case "append":
			found, o := slashrouter.FindOption("channel", c.Data.Options[0].Options)
			if !found {
				log.Printf("Unable to find channel option")
				return
			}
			channelID := o.ChannelValue(c.Session).ID
			found, o = slashrouter.FindOption("messageid", c.Data.Options[0].Options)
			if !found {
				log.Printf("Unable to find messageID option")
				return
			}
			messageID := o.StringValue()
			m, err := c.ChannelMessage(channelID, messageID)
			if err != nil {
				log.Printf("Error finding message: %s", err)
				return
			}
			found, o = slashrouter.FindOption("text", c.Data.Options[0].Options)
			if !found {
				log.Printf("Unable to find text option")
				return
			}
			text := o.StringValue()
			title := m.Embeds[0].Title
			desc := m.Embeds[0].Description + "\n" + text
			color := m.Embeds[0].Color
			author := m.Embeds[0].Author
			thumbnail := m.Embeds[0].Thumbnail
			footer := m.Embeds[0].Footer
			timestamp := m.Embeds[0].Timestamp
			if _, err := c.ChannelMessageEditEmbed(channelID, messageID, &discordgo.MessageEmbed{
				Title:       title,
				Description: desc,
				Color:       color,
				Author:      author,
				Thumbnail:   thumbnail,
				Footer:      footer,
				Timestamp:   timestamp,
			}); err != nil {
				log.Printf("Error sending message: %s", err)
			}
			err = c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			})
			if err != nil {
				log.Printf("Interaction response error: %s", err)
			}
		}

	},
	"avatar": func(c *slashrouter.Context) {
		var u *discordgo.User
		if found, o := slashrouter.FindOption("user", c.Data.Options); found {
			u = o.UserValue(c.Session)
		}
		err := c.InteractionRespond(c.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
		if err != nil {
			log.Printf("Error sending message: %s", err)
		}
		c.ChannelMessageSendEmbed(c.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    u.String(),
				IconURL: u.AvatarURL(""),
			},
			Image: &discordgo.MessageEmbedImage{
				URL: u.AvatarURL("2048"),
			},
		})
	},
}
