package growinglevelhandlers

import (
	"net/http"
	growinglevelcontrol "totally-legit-grow-management/v1/internal/logic/growingLevelsControl"
)

var _ IGrowingLevelHandler = &GrowingLevel{}

type IGrowingLevelHandler interface {
	CreateGrowingLevel(w http.ResponseWriter, r *http.Request)
	DeleteGrowingLevel(w http.ResponseWriter, r *http.Request)
	EditGrowingLevel(w http.ResponseWriter, r *http.Request)
	GetGrowingLevel(w http.ResponseWriter, r *http.Request)
	GetAllGrowingLevels(w http.ResponseWriter, r *http.Request)
}

type GrowingLevel struct {
	Logic growinglevelcontrol.IGrowingLevelsLogic
}
