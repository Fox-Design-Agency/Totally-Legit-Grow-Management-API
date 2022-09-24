// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowingMediumsLogic = &GrowingMediumControl{}

//
type IGrowingMediumsLogic interface {
	CreateGrowingMedium(*routemodels.CreateGrowingMediumRequest) (*routemodels.CreateGrowingMediumResponse, error)
	DeleteGrowingMedium(*routemodels.DeleteGrowingMediumRequest) error
	EditGrowingMedium(*routemodels.EditGrowingMediumRequest) (*routemodels.EditGrowingMediumResponse, error)
	GetGrowingMedium(*routemodels.GetGrowingMediumRequest) (*routemodels.GetGrowingMediumResponse, error)
	GetAllGrowingMediums(*routemodels.GetAllGrowingMediumsRequest) (*routemodels.GetAllGrowingMediumsResponse, error)
}

type GrowingMediumControl struct {
	Persistence persistence.IGrowingMediumsDB
}
