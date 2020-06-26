package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE super (
				id SERIAL PRIMARY KEY,
				uuid VARCHAR(40) NOT NULL UNIQUE,
				name VARCHAR(100) NOT NULL,
				full_name VARCHAR(255) NOT NULL,
				intelligence INT,
				power INT,
				occupation VARCHAR(100),
				image VARCHAR(100) UNIQUE,
				groups TEXT,
				category VARCHAR(100),
				number_relatives INT
			)
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE super")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200626133003_create_super_table", up, down, opts)
}
