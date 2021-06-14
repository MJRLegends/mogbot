package cache

import (
	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
)

type Cache struct {
	*GuildService
	*MemberService
}

func New(size int, db mogbot.Database) Cache {
	return Cache{
		GuildService:  NewGuildService(size, db),
		MemberService: NewMemberService(size, db),
	}
}
