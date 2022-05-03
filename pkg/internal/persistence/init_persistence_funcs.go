package persistence

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // should be here
)

func NewPersistence(dialect, connectionInfo string) (*Persistence, error) {
	db, err := sqlx.Open(dialect, connectionInfo)
	if err != nil {
		log.Println("Error connecting to DB")
		log.Println(err)
		return nil, err
	}

	return &Persistence{
		Postgres: db,
	}, nil
}
