// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package seedcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence/seeds"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ ISeedsLogic = &SeedControl{}

//
type ISeedsLogic interface {
	CreateSeed(*routemodels.CreateSeedRequest) (*routemodels.CreateSeedResponse, error)
	DeleteSeed(*routemodels.DeleteSeedRequest) error
	EditSeed(*routemodels.EditSeedRequest) (*routemodels.EditSeedResponse, error)
	GetSeed(*routemodels.GetSeedRequest) (*routemodels.GetSeedResponse, error)
	GetAllSeeds(*routemodels.GetAllSeedsRequest) (*routemodels.GetAllSeedsResponse, error)
}

type SeedControl struct {
	Persistence seeds.ISeedsDB
}
