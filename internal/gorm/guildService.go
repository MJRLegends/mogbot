package gorm

import (
	"log"

	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"gorm.io/gorm"
)

type GuildService struct {
	*gorm.DB
}

func (s *GuildService) AddGuild(g *mogbot.Guild) error {
	if r := s.Create(&g); r.Error != nil {
		log.Printf("Error adding %v to database: %s", g, r.Error)
		return r.Error
	}
	return nil
}

func (s *GuildService) GetGuild(gid string) (*mogbot.Guild, error) {
	var g *mogbot.Guild
	r := s.Where(&mogbot.Guild{ID: gid}).Take(g)
	if r.Error != nil {
		log.Printf("Error getting guild %s: %s", gid, r.Error)
		return nil, r.Error
	}
	return g, nil
}

func (s *GuildService) GetAllGuilds() ([]mogbot.Guild, error) {
	var guilds []mogbot.Guild
	r := s.Find(&guilds)
	if r.Error != nil {
		log.Printf("Error getting guild guilds %s", r.Error)
		return nil, r.Error
	}
	return guilds, nil
}

func (s *GuildService) UpdateGuild(guildID string, fields map[string]interface{}) error {
	s.Where(&mogbot.Guild{ID: guildID}).Updates(fields)
	return nil
}

func (s *GuildService) RemoveGuild(guildID string) error {
	if r := s.Delete(&mogbot.Guild{ID: guildID}); r.Error != nil {
		log.Printf("Error removing guild %s: %s", guildID, r.Error)
		return r.Error
	}
	return nil
}
