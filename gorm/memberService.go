package gorm

import (
	"log"

	mogbot "github.com/ChrisMcDearman/mogbot/mogbot"

	"gorm.io/gorm"
)

type MemberService struct {
	*gorm.DB
}

func (s *MemberService) AddMember(m *mogbot.Member) error {
	if r := s.Create(&m); r.Error != nil {
		log.Printf("Error adding %v to database: %s", m, r.Error)
		return r.Error
	}
	return nil
}

func (s *MemberService) GetMember(uid, gid string) (*mogbot.Member, error) {
	var m *mogbot.Member
	r := s.Where(&mogbot.Member{UserID: uid, GuildID: gid}).Take(m)
	if r.Error != nil {
		log.Printf("Error getting member with UserID='%s'in guild '%s': %s", uid, gid, r.Error)
		return nil, r.Error
	}
	return m, nil
}

func (s *MemberService) GetGuildMembers(gid string) ([]*mogbot.Member, error) {
	var members []*mogbot.Member
	r := s.Where(&mogbot.Member{GuildID: gid}).Find(&members)
	if r.Error != nil {
		log.Printf("Error getting guild '%s' members %s", gid, r.Error)
		return nil, r.Error
	}
	return members, nil
}

func (s *MemberService) UpdateMember(userID, guildID string, fields map[string]interface{}) (*mogbot.Member, error) {
	m := &mogbot.Member{UserID: userID, GuildID: guildID}
	if r := s.Model(m).Updates(fields); r.Error != nil {
		log.Printf("Error updating user '%s' in guild '%s': %s", userID, guildID, r.Error)
		return m, r.Error
	}
	return m, nil
}

func (s *MemberService) RemoveMember(uid, gid string) error {
	if r := s.Delete(&mogbot.Member{UserID: uid, GuildID: gid}); r.Error != nil {
		log.Printf("Error removing member with UserID='%s' in guild '%s': %s", uid, gid, r.Error)
		return r.Error
	}
	return nil
}
