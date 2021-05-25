package cache

import "github.com/captainmog/mogbot/internal/mogbot"

type Cache struct {
	*MemberService
}

func NewCache(size int, db mogbot.Database) Cache {
	return Cache{
		MemberService: NewMemberService(size, db),
	}

}
