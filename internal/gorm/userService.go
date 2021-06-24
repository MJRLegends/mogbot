package gorm

import (
	"log"

	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"gorm.io/gorm"
)

type UserService struct {
	*gorm.DB
}

func (s *UserService) AddUser(u *mogbot.User) error {
	if r := s.Create(&u); r.Error != nil {
		log.Printf("Error adding %v to database: %s", u, r.Error)
		return r.Error
	}
	return nil
}

func (s *UserService) GetUser(userID string) (*mogbot.User, error) {
	var u *mogbot.User
	r := s.Where(&mogbot.User{ID: userID}).Take(u)
	if r.Error != nil {
		log.Printf("Error getting user '%s': %s", userID, r.Error)
		return nil, r.Error
	}
	return u, nil
}

func (s *UserService) GetAllUsers() ([]*mogbot.User, error) {
	var users []*mogbot.User
	r := s.Find(&users)
	if r.Error != nil {
		log.Printf("Error getting all users: %s", r.Error)
		return nil, r.Error
	}
	return users, nil
}

func (s *UserService) UpdateUser(userID string, fields map[string]interface{}) (*mogbot.User, error) {
	u := &mogbot.User{ID: userID}
	if r := s.Model(u).Updates(fields); r.Error != nil {
		log.Printf("Error updating user '%s': %s", userID, r.Error)
		return u, r.Error
	}
	return u, nil
}

func (s *UserService) RemoveUser(userID string) error {
	if r := s.Delete(&mogbot.User{ID: userID}); r.Error != nil {
		log.Printf("Error removing member with UserID='%s': %s", userID, r.Error)
		return r.Error
	}
	return nil
}
