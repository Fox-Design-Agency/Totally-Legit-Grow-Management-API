// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func growSpotPlantsRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/growing-spot-plant", svr.GrowSpotPlants.CreateGrowSpotPlant).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.GrowSpotPlants.DeleteGrowSpotPlant).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.GrowSpotPlants.EditGrowSpotPlant).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-spot-plant", svr.GrowSpotPlants.GetGrowSpotPlant).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-spot-plants", svr.GrowSpotPlants.GetAllGrowSpotPlants).Methods(http.MethodGet)
}
