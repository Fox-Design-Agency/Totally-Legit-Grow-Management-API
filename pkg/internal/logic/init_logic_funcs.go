package logic

import (
	"log"
	"totally-legit-grow-management/v1/pkg/internal/persistence"
	"totally-legit-grow-management/v1/resources/config"
)

func NewLogicControl(cfg config.Config) (*Logic, error) {
	db, err := persistence.NewPersistence(cfg.Database.Dialect(), cfg.Database.Connection())
	if err != nil {
		log.Fatal("Couldn't Make DB Connection")
	}

	return &Logic{
		Persistence: db,
	}, nil
}
