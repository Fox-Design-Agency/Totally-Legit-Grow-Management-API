// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func plantStageRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/plant-stage", svr.CreatePlantStage).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/care", svr.CreatePlantStageCare).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.CreatePlantStageNutrients).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/connect", svr.ConnectPlantStage).Methods(http.MethodPost)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.DeletePlantStage).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-stage", svr.EditPlantStage).Methods(http.MethodPut)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.GetPlantStageByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/care", svr.GetPlantStageCareByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.GetPlantStageNutrient).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stages", svr.GetAllPlantStages).Methods(http.MethodGet)
}