// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence/plants"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IPlantsLogic = &PlantControl{}

//
type IPlantsLogic interface {
	CreatePlant(*routemodels.CreatePlantRequest) (*routemodels.CreatePlantResponse, error)
	DeletePlant(*routemodels.DeletePlantRequest) error
	EditPlant(*routemodels.EditPlantRequest) (*routemodels.EditPlantResponse, error)
	GetPlant(*routemodels.GetPlantRequest) (*routemodels.GetPlantResponse, error)
	GetAllPlants(*routemodels.GetAllPlantsRequest) (*routemodels.GetAllPlantsResponse, error)
}

type PlantControl struct {
	Persistence plants.IPlantsDB
}
