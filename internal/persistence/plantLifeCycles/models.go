// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecycle

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantLifeCyclesDB = &PlantLifecyclePersistence{}

//
type IPlantLifeCyclesDB interface {
	CreatePlantLifeCycle(*routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	CreatePlantLifeCycleWithTransaction(*sqlx.Tx, *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	DeletePlantLifeCycle(*routemodels.DeletePlantLifeCycleRequest) error
	EditPlantLifeCycle(*routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error)
	GetPlantLifeCycleByID(*routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetPlantLifeCycleByIDWithTransaction(*sqlx.Tx, *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetAllPlantLifeCycles(*routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
	GetAllPlantLifeCyclesWithTransaction(*sqlx.Tx, *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
}

type PlantLifecyclePersistence struct {
	Postgres *sqlx.DB
}
