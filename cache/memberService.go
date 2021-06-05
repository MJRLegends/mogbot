package cache

import (
	"log"

	"github.com/pkg/errors"

	"github.com/ChrisMcDearman/mogbot/mogbot"

	lru "github.com/hashicorp/golang-lru"
)

type MemberService struct {
	*lru.Cache
	mogbot.MemberService
}

func NewMemberService(size int, ms mogbot.MemberService) *MemberService {
	l, err := lru.New(size)
	if err != nil {
		log.Printf("Error creating new cache")
		return nil
	}
	return &MemberService{Cache: l, MemberService: ms}
}

func (s *MemberService) AddMember(m *mogbot.Member) error {
	err := s.MemberService.AddMember(m)
	if err != nil {
		return err
	}
	s.Add(m.UserID+m.GuildID, m)
	log.Printf("Added %v to cache", m)
	return nil
}

func (s *MemberService) GetMember(userID, guildID string) (*mogbot.Member, error) {
	user, ok := s.Get(userID + guildID)
	if !ok {
		m, err := s.MemberService.GetMember(userID, guildID)
		if err != nil {
			return nil, err
		}
		s.Add(m.UserID+m.GuildID, *m)
		return m, nil
	}
	m := user.(mogbot.Member)
	log.Printf("Retrieved %v from cache", m)
	return &m, nil
}

func (s *MemberService) UpdateMember(userID, guildID string, fields map[string]interface{}) error {
	s.Remove(userID + guildID)
	v, ok := s.Get(userID + guildID)
	if !ok {
		return errors.New("")
	}
	t, _ := v.(*mogbot.Member)
	if err := t.FillStruct(fields); err != nil {
		return err
	}
	return s.MemberService.UpdateMember(userID, guildID, fields)
}

func (s *MemberService) RemoveMember(userID, guildID string) error {
	s.Remove(userID + guildID)
	return s.MemberService.RemoveMember(userID, guildID)
}
