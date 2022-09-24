// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tasks

import "github.com/jmoiron/sqlx"

func InitTaskPersistence(db *sqlx.DB) *TaskPersistence {
	return &TaskPersistence{
		Postgres: db,
	}
}
