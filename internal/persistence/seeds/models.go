// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package seeds

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ ISeedsDB = &SeedPersistence{}

//
type ISeedsDB interface {
	CreateSeed(*routemodels.CreateSeedRequest) (*routemodels.CreateSeedResponse, error)
	DeleteSeed(*routemodels.DeleteSeedRequest) error
	EditSeed(*routemodels.EditSeedRequest) (*routemodels.EditSeedResponse, error)
	GetSeed(*routemodels.GetSeedRequest) (*routemodels.GetSeedResponse, error)
	GetAllSeeds(*routemodels.GetAllSeedsRequest) (*routemodels.GetAllSeedsResponse, error)
}
type SeedPersistence struct {
	Postgres *sqlx.DB
}
