// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocations

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingLocationsDB = &GrowingLocationPersistence{}

//
type IGrowingLocationsDB interface {
	CreateGrowingLocation(*routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	CreateGrowingLocationWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	DeleteGrowingLocation(*routemodels.DeleteGrowingLocationRequest) error
	EditGrowingLocation(*routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error)
	GetGrowingLocationByID(*routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetGrowingLocationByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetAllGrowingLocationsByGrowingGroupID(*routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
	GetAllGrowingLocationsByGrowingGroupIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
}

type GrowingLocationPersistence struct {
	Postgres *sqlx.DB
}
