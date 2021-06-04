package cache

import "github.com/ChrisMcDearman/mogbot/mogbot"

type Cache struct {
	*MemberService
}

func New(size int, db mogbot.Database) Cache {
	return Cache{
		MemberService: NewMemberService(size, db),
	}
}
