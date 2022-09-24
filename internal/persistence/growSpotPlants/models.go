// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplants

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowSpotPlantsDB = &GrowSpotPlantsPersistence{}

//
type IGrowSpotPlantsDB interface {
	CreateGrowSpotPlant(*routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error)
	DeleteGrowSpotPlant(*routemodels.DeleteGrowSpotPlantRequest) error
	EditGrowSpotPlant(*routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error)
	GetGrowSpotPlant(*routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error)
	GetAllGrowSpotPlants(*routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error)
}

type GrowSpotPlantsPersistence struct {
	Postgres *sqlx.DB
}
