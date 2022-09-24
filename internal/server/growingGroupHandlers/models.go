package growinggrouphandlers

import (
	"net/http"
	growinggroupcontrol "totally-legit-grow-management/v1/internal/logic/growingGroupsControl"
)

var _ IGrowingGroupHandler = &GrowingGroup{}

type IGrowingGroupHandler interface {
	CreateGrowingGroup(w http.ResponseWriter, r *http.Request)
	DeleteGrowingGroup(w http.ResponseWriter, r *http.Request)
	EditGrowingGroup(w http.ResponseWriter, r *http.Request)
	GetGrowingGroup(w http.ResponseWriter, r *http.Request)
	GetAllGrowingGroups(w http.ResponseWriter, r *http.Request)
}

type GrowingGroup struct {
	Logic growinggroupcontrol.IGrowingGroupsLogic
}
