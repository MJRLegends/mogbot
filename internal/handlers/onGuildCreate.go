package handlers

import (
	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"github.com/bwmarrin/discordgo"
)

func OnGuildCreate(b *mogbot.Bot) interface{} {
	return func(s *discordgo.Session, g *discordgo.GuildCreate) {
		//for _, r := range g.Roles {
		//	if err := b.DB.AddRole(*bot.NewRole(r)); err != nil {
		//		continue
		//	}
		//}
		//roles, err := b.DB.GetAllRoles()
		//if err != nil {
		//	panic(err)
		//}
		//for _, m := range g.Members {
		//	nm := &mogbot.Member{m.User.ID, m.GuildID, []mogbot.Role{}}
		//	for _, r := range m.Roles {
		//		if fr, found := mogbot.FindRole(r, roles); found {
		//			nm.Roles = append(nm.Roles, *fr)
		//		}
		//	}
		//	if err := b.DB.AddMember(nm); err != nil {
		//		if err = b.DB.UpdateMember(nm); err != nil {
		//			continue
		//		}
		//	}
		//}
	}
}
