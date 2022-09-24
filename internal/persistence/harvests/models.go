// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package harvests

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IHarvestsDB = &HarvestPersistence{}

//
type IHarvestsDB interface {
	CreateHarvest(*routemodels.CreateHarvestRequest) (*routemodels.CreateHarvestResponse, error)
	DeleteHarvest(*routemodels.DeleteHarvestRequest) error
	EditHarvest(*routemodels.EditHarvestRequest) (*routemodels.EditHarvestResponse, error)
	GetHarvest(*routemodels.GetHarvestRequest) (*routemodels.GetHarvestResponse, error)
	GetAllHarvests(*routemodels.GetAllHarvestsRequest) (*routemodels.GetAllHarvestsResponse, error)
}

type HarvestPersistence struct {
	Postgres *sqlx.DB
}
