// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplanthandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IGrowSpotPlantsHandlers interface {
	CreateGrowSpotPlant(w http.ResponseWriter, r *http.Request)
	DeleteGrowSpotPlant(w http.ResponseWriter, r *http.Request)
	EditGrowSpotPlant(w http.ResponseWriter, r *http.Request)
	GetGrowSpotPlant(w http.ResponseWriter, r *http.Request)
	GetAllGrowSpotPlants(w http.ResponseWriter, r *http.Request)
}

type GrowSpotPlants struct {
	Logic logic.IGrowSpotPlantsLogic
}
