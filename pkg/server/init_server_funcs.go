package server

import (
	"log"

	"smart-grow-management-api/v1/pkg/internal/persistence"
	"smart-grow-management-api/v1/resources/config"
)

func NewServer(cfg config.Config) *Server {

	db, err := persistence.NewPersistence(cfg.Database.Dialect(), cfg.Database.Connection())
	if err != nil {
		log.Fatal("Couldn't Make DB Connection")
	}

	return &Server{
		Persistence: db,
	}
}
