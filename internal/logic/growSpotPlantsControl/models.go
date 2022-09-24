// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplantscontrol

import (
	growspotplants "totally-legit-grow-management/v1/internal/persistence/growSpotPlants"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowSpotPlantsLogic = &GrowSpotPlantControl{}

//
type IGrowSpotPlantsLogic interface {
	CreateGrowSpotPlant(*routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error)
	DeleteGrowSpotPlant(*routemodels.DeleteGrowSpotPlantRequest) error
	EditGrowSpotPlant(*routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error)
	GetGrowSpotPlant(*routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error)
	GetAllGrowSpotPlants(*routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error)
}

type GrowSpotPlantControl struct {
	Persistence growspotplants.IGrowSpotPlantsDB
}
