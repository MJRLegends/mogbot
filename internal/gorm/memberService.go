package gorm

import (
	"log"

	"github.com/captainmog/mogbot/internal/mogbot"
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
	r := s.Take(&mogbot.Member{UserID: uid, GuildID: gid})
	if r.Error != nil {
		log.Printf("Error getting member with UserID='%s'in guild '%s': %s", uid, gid, r.Error)
		return nil, r.Error
	}
	var m *mogbot.Member
	r.Scan(m)
	return m, nil
}

func (s *MemberService) GetGuildMembers(gid string) ([]*mogbot.Member, error) {
	r := s.Find(&mogbot.Member{GuildID: gid})
	if r.Error != nil {
		log.Printf("Error getting guild '%s' members %s", gid, r.Error)
		return nil, r.Error
	}
	var m []*mogbot.Member
	r.Scan(m)
	return m, nil
}

func (s *MemberService) UpdateMember(nm *mogbot.Member) error {
	var dbm *mogbot.Member
	s.First(dbm)
	dbm.Roles = nm.Roles
	s.Save(dbm)
	return nil
}

func (s *MemberService) RemoveMember(uid, gid string) error {
	if r := s.Delete(&mogbot.Member{UserID: uid, GuildID: gid}); r.Error != nil {
		log.Printf("Error removing member with UserID='%s' in guild '%s': %s", uid, gid, r.Error)
		return r.Error
	}
	return nil
}
