// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func plantStageRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/plant-stage", svr.PlantStage.CreatePlantStage).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/care", svr.PlantStage.CreatePlantStageCare).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.PlantStage.CreatePlantStageNutrients).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-stage/connect", svr.PlantStage.ConnectPlantStage).Methods(http.MethodPost)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.PlantStage.DeletePlantStage).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-stage", svr.PlantStage.EditPlantStage).Methods(http.MethodPut)
	//
	r.HandleFunc("/api/v1/plant-stage", svr.PlantStage.GetPlantStageByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/care", svr.PlantStage.GetPlantStageCareByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stage/nutrient", svr.PlantStage.GetPlantStageNutrient).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-stages", svr.PlantStage.GetAllPlantStages).Methods(http.MethodGet)
}
