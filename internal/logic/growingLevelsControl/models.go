// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglevelcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowingLevelsLogic = &GrowingLevelControl{}

//
type IGrowingLevelsLogic interface {
	CreateGrowingLevel(*routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	DeleteGrowingLevel(*routemodels.DeleteGrowingLevelRequest) error
	EditGrowingLevel(*routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error)
	GetGrowingLevelByID(*routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetAllGrowingLevelsByGrowingLocationID(*routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
}

type GrowingLevelControl struct {
	Persistence persistence.IGrowingLevelsDB
}
