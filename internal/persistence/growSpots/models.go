// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspots

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowSpotsDB = &GrowSpotsPersistence{}

//
type IGrowSpotsDB interface {
	CreateGrowSpot(*routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	CreateGrowSpotWithTransaction(*sqlx.Tx, *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	DeleteGrowSpot(*routemodels.DeleteGrowSpotRequest) error
	EditGrowSpot(*routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error)
	GetGrowSpotByID(*routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetGrowSpotByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetAllGrowSpotsByGrowLevelID(*routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
	GetAllGrowSpotsByGrowLevelIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
}

type GrowSpotsPersistence struct {
	Postgres *sqlx.DB
}
