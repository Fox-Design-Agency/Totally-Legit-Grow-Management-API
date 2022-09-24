// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantstagescontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IPLantStagesLogic = &PlantStageControl{}

//
type IPLantStagesLogic interface {
	CreatePlantStage(*routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageCare(*routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageNutrients(*routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	ConnectPlantStage(*routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	DeletePlantStage(*routemodels.DeletePlantStageRequest) error
	EditPlantStage(*routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error)
	//
	GetPlantStageByID(*routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageCareByID(*routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error)
	GetPlantStageNutrient(*routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error)
	GetAllPlantStages(*routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
}

type PlantStageControl struct {
	Persistence persistence.IPlantStagesDB
}
