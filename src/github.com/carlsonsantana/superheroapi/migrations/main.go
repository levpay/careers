package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

const directory = "."

func main() {
	db := pg.Connect(&pg.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
		),
		User:     os.Getenv("DATABASE_USER"),
		Database: os.Getenv("DATABASE_NAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	})

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
