package router

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func FindOption(name string, options []*discordgo.ApplicationCommandInteractionDataOption) (bool, *discordgo.ApplicationCommandInteractionDataOption) {
	for _, o := range options {
		if o.Name == name {
			return true, o
		}
	}
	return false, nil
}

func FindCommand(name string, cmds []*discordgo.ApplicationCommand) bool {
	for _, c := range cmds {
		if c.Name == name {
			return true
		}
	}
	return false
}

//func ContainsCommand(cmds []*discordgo.ApplicationCommand, c *discordgo.ApplicationCommand) bool {
//	for _, cmd := range cmds {
//		if cmp.Equal(*cmd, *c) {
//			return true
//		}
//	}
//	return false
//}

func WaitForMessage(s *discordgo.Session) chan *discordgo.MessageCreate {
	channel := make(chan *discordgo.MessageCreate)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		channel <- e
	})
	return channel
}

func WaitForUserMessage(s *discordgo.Session, userID string) chan *discordgo.MessageCreate {
	channel := make(chan *discordgo.MessageCreate)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		if e.Author.ID == userID {
			channel <- e
		}
	})
	return channel
}

func WaitForUserReact(s *discordgo.Session, userID string) chan *discordgo.MessageReactionAdd {
	channel := make(chan *discordgo.MessageReactionAdd)
	s.AddHandler(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		if e.UserID == userID {
			channel <- e
		}
	})
	return channel
}

func HasPerm(userPerms, hasPerm int64) bool {
	return userPerms&hasPerm != 0
}

func IsChannelAdmin(s *discordgo.Session, channelID, userID string) bool {
	p, err := s.UserChannelPermissions(userID, channelID)
	if err != nil {
		log.Printf("Error retrieving user channel permissions: %s", err)
		return false
	}
	if HasPerm(p, discordgo.PermissionAdministrator) {
		return true
	}
	return false
}

//
//func CanBan(s *discordgo.Session, guildID string, m *discordgo.Member) bool {
//	for _, r := range m.Roles {
//		role, err := s.State.Role(guildID, r)
//		if err != nil {
//			log.Printf("Error retrieving '%s' role: %s", r, err)
//			return false
//		}
//		if HasPerm(role, discordgo.PermissionBanMembers) || HasPerm(role, discordgo.PermissionAdministrator) {
//			return true
//		}
//	}
//	return false
//}
