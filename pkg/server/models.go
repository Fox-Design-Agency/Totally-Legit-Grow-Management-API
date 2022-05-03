package server

import (
	"totally-legit-grow-management/v1/pkg/internal/persistence"
)

type Server struct {
	Persistence *persistence.Persistence
}
