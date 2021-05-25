package postgres

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
	*MemberService
}

func LoadSchema(mode string) (*DB, error) {
	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		return nil, err
	}
	db, err := NewConnection(mode)
	if err != nil {
		return nil, err
	}
	db.DB.MustExec(string(schema))
	return db, nil
}

func NewConnection(mode string) (*DB, error) {
	timeoutCd, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	db, err := sqlx.ConnectContext(timeoutCd, "postgres", os.Getenv(mode))
	if err != nil {
		log.Printf("Error connecting to postgres db: %s", err)
		return nil, err
	}
	return &DB{
		db,
		&MemberService{db},
	}, nil
}
