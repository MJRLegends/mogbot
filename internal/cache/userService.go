package cache

import (
	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	lru "github.com/hashicorp/golang-lru"
	"log"
)

type UserService struct {
	*lru.Cache
	mogbot.UserService
}

func NewUserService(size int, ms mogbot.UserService) *UserService {
	l, err := lru.New(size)
	if err != nil {
		log.Printf("Error creating new cache")
		return nil
	}
	return &UserService{Cache: l, UserService: ms}
}

func (s *UserService) AddUser(u *mogbot.User) error {
	err := s.UserService.AddUser(u)
	if err != nil {
		return err
	}
	s.Add(u.ID, u)
	log.Printf("Added user '%s' to cache", u.ID)
	return nil
}

func (s *UserService) GetUser(userID string) (*mogbot.User, error) {
	User, ok := s.Get(userID)
	if !ok {
		u, err := s.UserService.GetUser(userID)
		if err != nil {
			return nil, err
		}
		s.Add(userID, u)
		return u, nil
	}
	u := User.(*mogbot.User)
	log.Printf("Retrieved user '%s' from cache", u.ID)
	return u, nil
}

func (s *UserService) UpdateUser(userID string, fields map[string]interface{}) (*mogbot.User, error) {
	u, err := s.UserService.UpdateUser(userID, fields)
	s.Remove(userID)
	s.Add(userID, u)
	return u, err
}

func (s *UserService) RemoveUser(userID string) error {
	s.Remove(userID)
	return s.UserService.RemoveUser(userID)
}
