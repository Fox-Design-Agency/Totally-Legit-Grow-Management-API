// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmedium

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingMediumsDB = &GrowingMediumPersistence{}

//
type IGrowingMediumsDB interface {
	CreateGrowingMedium(*routemodels.CreateGrowingMediumRequest) (*routemodels.CreateGrowingMediumResponse, error)
	DeleteGrowingMedium(*routemodels.DeleteGrowingMediumRequest) error
	EditGrowingMedium(*routemodels.EditGrowingMediumRequest) (*routemodels.EditGrowingMediumResponse, error)
	GetGrowingMedium(*routemodels.GetGrowingMediumRequest) (*routemodels.GetGrowingMediumResponse, error)
	GetAllGrowingMediums(*routemodels.GetAllGrowingMediumsRequest) (*routemodels.GetAllGrowingMediumsResponse, error)
}

type GrowingMediumPersistence struct {
	Postgres *sqlx.DB
}
