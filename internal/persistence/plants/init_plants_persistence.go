// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plants

import "github.com/jmoiron/sqlx"

func InitPlantPersistence(db *sqlx.DB) *PlantPersistence {
	return &PlantPersistence{
		Postgres: db,
	}
}
