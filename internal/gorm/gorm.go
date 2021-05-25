package gorm

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	*MemberService
}

func NewConnection(mode string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv(mode)), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to postgres db: %s", err)
		return nil, err
	}
	return &DB{
		db,
		&MemberService{db},
	}, nil
}
