package gorm

import (
	"log"
	"os"

	"github.com/ChrisMcDearman/mogbot/internal/mogbot"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*UserService
	//*GuildService
	//*MemberService
}

func NewConnection(source string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv(source)), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to postgres db: %s", err)
		return nil, err
	}
	db.AutoMigrate(
		&mogbot.User{},
		//&mogbot.Guild{},
		//&mogbot.Member{},
		//&mogbot.Role{},
		//&mogbot.Log{},
	)
	return &Database{
		&UserService{db},
		//&GuildService{db},
		//&MemberService{db},
	}, nil
}
