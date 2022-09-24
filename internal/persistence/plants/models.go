// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plants

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantsDB = &PlantPersistence{}

//
type IPlantsDB interface {
	CreatePlant(*routemodels.CreatePlantRequest) (*routemodels.CreatePlantResponse, error)
	DeletePlant(*routemodels.DeletePlantRequest) error
	EditPlant(*routemodels.EditPlantRequest) (*routemodels.EditPlantResponse, error)
	GetPlant(*routemodels.GetPlantRequest) (*routemodels.GetPlantResponse, error)
	GetAllPlants(*routemodels.GetAllPlantsRequest) (*routemodels.GetAllPlantsResponse, error)
}

type PlantPersistence struct {
	Postgres *sqlx.DB
}
