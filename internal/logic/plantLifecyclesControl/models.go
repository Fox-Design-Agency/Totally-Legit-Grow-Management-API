// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecyclecontrol

import (
	plantlifecycle "totally-legit-grow-management/v1/internal/persistence/plantLifeCycles"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IPlantLifeCyclesLogic = &PlantLifeCycleControl{}

//
type IPlantLifeCyclesLogic interface {
	CreatePlantLifeCycle(*routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error)
	DeletePlantLifeCycle(*routemodels.DeletePlantLifeCycleRequest) error
	EditPlantLifeCycle(*routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error)
	GetPlantLifeCycleByID(*routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error)
	GetAllPlantLifeCycles(*routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error)
}

type PlantLifeCycleControl struct {
	Persistence plantlifecycle.IPlantLifeCyclesDB
}
