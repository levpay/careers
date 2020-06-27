package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			ALTER TABLE super ADD COLUMN superheroapi_id INT UNIQUE;
			UPDATE super SET superheroapi_id = -id;
			ALTER TABLE super ALTER COLUMN superheroapi_id SET NOT NULL;
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("ALTER TABLE super DROP COLUMN superheroapi_id")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200626181422_add_superheroapi_id_column", up, down, opts)
}
