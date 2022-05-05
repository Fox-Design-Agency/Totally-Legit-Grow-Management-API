package server

import (
	"log"

	"totally-legit-grow-management/v1/pkg/internal/logic"
	"totally-legit-grow-management/v1/resources/config"
)

func NewServer(cfg config.Config, shouldMigrate bool) *Server {

	control, err := logic.NewLogicControl(cfg)
	if err != nil {
		log.Fatal("Logic Layer Not Setup!")
	}

	if shouldMigrate {
		err := control.Persistence.MigrateDBUP()
		if err != nil {
			// migrations couldn't happen
			log.Println(err)
		}
	}

	return &Server{
		Logic: control,
	}
}
