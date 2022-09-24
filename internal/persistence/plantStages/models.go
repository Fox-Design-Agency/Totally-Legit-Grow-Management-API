// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantstages

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantStagesDB = &PlantStagesPersistence{}

//
type IPlantStagesDB interface {
	CreatePlantStage(*routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error)
	CreatePlantStageCare(*routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageCareWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error)
	CreatePlantStageNutrients(*routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	CreatePlantStageNutrientsWithTransaction(*sqlx.Tx, *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error)
	ConnectPlantStage(*routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	ConnectPlantStageWithTransaction(*sqlx.Tx, *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error)
	//
	DeletePlantStage(*routemodels.DeletePlantStageRequest) error
	EditPlantStage(*routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error)
	//
	GetPlantStageByID(*routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageByIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error)
	GetPlantStageCareByID(*routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error)
	GetPlantStageCareByPlantStageIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantStageCareByPlantStageIDRequest) (*routemodels.GetPlantStageCareByPlantStageIDResponse, error)
	GetPlantStageNutrient(*routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error)
	//
	GetAllPlantStages(*routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
	GetAllPlantStagesWithTransaction(*sqlx.Tx, *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error)
}

type PlantStagesPersistence struct {
	Postgres *sqlx.DB
}
