// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotcontrol

import (
	growspots "totally-legit-grow-management/v1/internal/persistence/growSpots"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowSpotsLogic = &GrowSpotControl{}

//
type IGrowSpotsLogic interface {
	CreateGrowSpot(*routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error)
	DeleteGrowSpot(*routemodels.DeleteGrowSpotRequest) error
	EditGrowSpot(*routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error)
	GetGrowSpotByID(*routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error)
	GetAllGrowSpotsByGrowLevelID(*routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error)
}

type GrowSpotControl struct {
	Persistence growspots.IGrowSpotsDB
}
