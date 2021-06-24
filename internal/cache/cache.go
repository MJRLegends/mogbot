package cache

import (
	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
)

type Cache struct {
	*UserService
	//*GuildService
	//*MemberService
}

func New(size int, db mogbot.Database) Cache {
	return Cache{
		UserService: NewUserService(size, db),
		//GuildService:  NewGuildService(size, db),
		//MemberService: NewMemberService(size, db),
	}
}
