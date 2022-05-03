package persistence

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// MigrateDBUP runs the migration files up
func (db *Persistence) MigrateDBUP() error {
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

//MigrateDBDown runs the migration files down
func (db *Persistence) MigrateDBDown() error {
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

func (db *Persistence) StartTransaction() (*sqlx.Tx, error) {
	tx, err := db.Postgres.Beginx()
	if err != nil {
		// log error
		return nil, err
	}
	return tx, nil
}
