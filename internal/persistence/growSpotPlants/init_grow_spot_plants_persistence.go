// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplants

import "github.com/jmoiron/sqlx"

func InitGrowSpotPlantsPersistence(db *sqlx.DB) *GrowSpotPlantsPersistence {
	return &GrowSpotPlantsPersistence{
		Postgres: db,
	}
}
