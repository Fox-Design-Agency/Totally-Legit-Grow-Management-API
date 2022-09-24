// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// MigrateDBUP runs the migration files found in the resources folder up
func (db *DBControl) MigrateDBUP() error {
	// run the migrate here
	migrations := &migrate.FileMigrationSource{
		Dir: "resources/migrations",
	}
	n, err := migrate.Exec(db.Postgres.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
		log.Println(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}

//MigrateDBDown runs the migration files found in the resources folder down
func (db *DBControl) MigrateDBDown() error {
	// run the migrate here
	migrations := &migrate.FileMigrationSource{
		Dir: "resources/migrations",
	}
	n, err := migrate.Exec(db.Postgres.DB, "postgres", migrations, migrate.Down)
	if err != nil {
		// Handle errors!
		log.Println(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}

// StartTransaction will begin a transaction for the logic layer
// to utilize and compose persistence layers calls that require
// a transaction. This transaction will need to be commited within
// the logic layer prior to a successful return.
func (db *DBControl) StartTransaction() (*sqlx.Tx, error) {
	tx, err := db.Postgres.Beginx()
	if err != nil {
		// log error
		return nil, err
	}
	return tx, nil
}
