package gorm

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*GuildService
	*MemberService
}

func NewConnection(source string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv(source)), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to postgres db: %s", err)
		return nil, err
	}
	return &DB{
		&GuildService{db},
		&MemberService{db},
	}, nil
}
