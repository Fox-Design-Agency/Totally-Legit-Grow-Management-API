// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // should be here
)

func NewPersistence(dialect, connectionInfo string) (*Persistence, error) {
	db, err := sqlx.Open(dialect, connectionInfo)
	if err != nil {
		log.Println("Error connecting to DB")
		log.Println(err)
		return nil, err
	}

	return &Persistence{
		Postgres: db,
	}, nil
}
