// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package harvests

import "github.com/jmoiron/sqlx"

func InitHarvestPersistence(db *sqlx.DB) *HarvestPersistence {
	return &HarvestPersistence{
		Postgres: db,
	}
}
