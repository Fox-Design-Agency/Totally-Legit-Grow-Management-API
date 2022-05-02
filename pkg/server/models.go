package server

import (
	"smart-grow-management-api/v1/pkg/internal/persistence"
)

type Server struct {
	Persistence *persistence.Persistence
}
