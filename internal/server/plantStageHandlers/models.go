// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantstagehandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IPlantStageHandler interface {
	CreatePlantStage(w http.ResponseWriter, r *http.Request)
	CreatePlantStageCare(w http.ResponseWriter, r *http.Request)
	CreatePlantStageNutrients(w http.ResponseWriter, r *http.Request)
	ConnectPlantStage(w http.ResponseWriter, r *http.Request)
	DeletePlantStage(w http.ResponseWriter, r *http.Request)
	EditPlantStage(w http.ResponseWriter, r *http.Request)
	GetPlantStageByID(w http.ResponseWriter, r *http.Request)
	GetPlantStageCareByID(w http.ResponseWriter, r *http.Request)
	GetPlantStageNutrient(w http.ResponseWriter, r *http.Request)
	GetAllPlantStages(w http.ResponseWriter, r *http.Request)
}

type PlantStage struct {
	Logic logic.IPLantStagesLogic
}
