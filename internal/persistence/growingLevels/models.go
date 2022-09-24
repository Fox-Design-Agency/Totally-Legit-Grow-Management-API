// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglevels

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingLevelsDB = &GrowingLevelPersistence{}

//
type IGrowingLevelsDB interface {
	CreateGrowingLevel(*routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	CreateGrowingLevelWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error)
	DeleteGrowingLevel(*routemodels.DeleteGrowingLevelRequest) error
	EditGrowingLevel(*routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error)
	GetGrowingLevelByID(*routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetGrowingLevelByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error)
	GetAllGrowingLevelsByGrowingLocationID(*routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
	GetAllGrowingLevelsByGrowingLocationIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error)
}

type GrowingLevelPersistence struct {
	Postgres *sqlx.DB
}
