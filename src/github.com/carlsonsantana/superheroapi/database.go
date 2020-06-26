package superheroapi

import (
	"fmt"
	"os"

	"github.com/arthurkushman/buildsqlx"
	_ "github.com/lib/pq"
)

func GetDatabaseConnection() *buildsqlx.DB {
	return buildsqlx.NewDb(
		buildsqlx.NewConnection(
			"postgres",
			fmt.Sprintf(
				"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASSWORD"),
				os.Getenv("DATABASE_SSLMODE"),
			),
		),
	)
}
