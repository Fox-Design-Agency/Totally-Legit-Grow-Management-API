// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package harvestcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IHarvestsLogic = &HarvestControl{}

//
type IHarvestsLogic interface {
	CreateHarvest(*routemodels.CreateHarvestRequest) (*routemodels.CreateHarvestResponse, error)
	DeleteHarvest(*routemodels.DeleteHarvestRequest) error
	EditHarvest(*routemodels.EditHarvestRequest) (*routemodels.EditHarvestResponse, error)
	GetHarvest(*routemodels.GetHarvestRequest) (*routemodels.GetHarvestResponse, error)
	GetAllHarvests(*routemodels.GetAllHarvestsRequest) (*routemodels.GetAllHarvestsResponse, error)
}

type HarvestControl struct {
	Persistence persistence.IHarvestsDB
}
