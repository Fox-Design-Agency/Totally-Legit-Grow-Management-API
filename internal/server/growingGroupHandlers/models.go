package growinggrouphandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IGrowingGroupHandler interface {
	CreateGrowingGroup(w http.ResponseWriter, r *http.Request)
	DeleteGrowingGroup(w http.ResponseWriter, r *http.Request)
	EditGrowingGroup(w http.ResponseWriter, r *http.Request)
	GetGrowingGroup(w http.ResponseWriter, r *http.Request)
	GetAllGrowingGroups(w http.ResponseWriter, r *http.Request)
}

type GrowingGroup struct {
	Logic logic.IGrowingGroupsLogic
}
