// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocationcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowingLocationsLogic = &GrowingLocationControl{}

//
type IGrowingLocationsLogic interface {
	CreateGrowingLocation(*routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error)
	DeleteGrowingLocation(*routemodels.DeleteGrowingLocationRequest) error
	EditGrowingLocation(*routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error)
	GetGrowingLocationByID(*routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error)
	GetAllGrowingLocationsByGrowingGroupID(*routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error)
}

type GrowingLocationControl struct {
	Persistence persistence.IGrowingLocationsDB
}
