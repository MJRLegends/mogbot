package cache

import (
	"log"

	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	lru "github.com/hashicorp/golang-lru"
)

type GuildService struct {
	*lru.Cache
	mogbot.GuildService
}

func NewGuildService(size int, ms mogbot.GuildService) *GuildService {
	l, err := lru.New(size)
	if err != nil {
		log.Printf("Error creating new cache")
		return nil
	}
	return &GuildService{Cache: l, GuildService: ms}
}

func (s *GuildService) AddGuild(g *mogbot.Guild) error {
	err := s.GuildService.AddGuild(g)
	if err != nil {
		return err
	}
	s.Add(g.ID, g)
	log.Printf("Added guild %v to cache", g)
	return nil
}

func (s *GuildService) GetGuild(id string) (*mogbot.Guild, error) {
	guild, ok := s.Get(id)
	if !ok {
		g, err := s.GuildService.GetGuild(id)
		if err != nil {
			return nil, err
		}
		s.Add(id, g)
		return g, nil
	}
	g := guild.(*mogbot.Guild)
	log.Printf("Retrieved guild %v from cache", g)
	return g, nil
}

func (s *GuildService) UpdateGuild(guildID string, fields map[string]interface{}) (*mogbot.Guild, error) {
	g, err := s.GuildService.UpdateGuild(guildID, fields)
	s.Remove(guildID)
	s.Add(guildID, g)
	return g, err
}

func (s *GuildService) RemoveGuild(guildID string) error {
	s.Remove(guildID)
	return s.GuildService.RemoveGuild(guildID)
}
