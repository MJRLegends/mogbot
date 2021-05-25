package msgrouter

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type Context struct {
	*discordgo.Session
	*discordgo.Message
	Vars map[string]interface{}
}

func (c *Context) Reply(text string) error {
	_, err := c.ChannelMessageSend(c.Message.ChannelID, text)
	if err != nil {
		log.Printf("Error replying to message: %s", err)
		return err
	}
	return nil
}

func (c *Context) ReplyEmbed(e *discordgo.MessageEmbed) error {
	_, err := c.ChannelMessageSendEmbed(c.Message.ChannelID, e)
	if err != nil {
		log.Printf("Error replying to message: %s", err)
		return err
	}
	return nil
}