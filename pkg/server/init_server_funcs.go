package server

import (
	"log"

	"totally-legit-grow-management/v1/pkg/internal/logic"
	"totally-legit-grow-management/v1/resources/config"
)

func NewServer(cfg config.Config) *Server {

	control, err := logic.NewLogicControl(cfg)
	if err != nil {
		log.Fatal("Logic Layer Not Setup!")
	}

	return &Server{
		Logic: control,
	}
}
