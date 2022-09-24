// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspothandlers

import (
	"net/http"
	growspotcontrol "totally-legit-grow-management/v1/internal/logic/growSpotsControl"
)

var _ IGrowSpotHandler = &GrowSpot{}

type IGrowSpotHandler interface {
	CreateGrowSpot(w http.ResponseWriter, r *http.Request)
	DeleteGrowSpot(w http.ResponseWriter, r *http.Request)
	EditGrowSpot(w http.ResponseWriter, r *http.Request)
	GetGrowSpot(w http.ResponseWriter, r *http.Request)
	GetAllGrowSpots(w http.ResponseWriter, r *http.Request)
}

type GrowSpot struct {
	Logic growspotcontrol.IGrowSpotsLogic
}
