package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/dvdscripter/superheroapi/storage"
	"github.com/dvdscripter/superheroapi/storage/sql"
	"github.com/gorilla/mux"
)

// App is our depency injection container, should contain only singletons
type App struct {
	storage storage.SuperStorage
	log     *log.Logger
	config  *Config
}

func main() {
	configPath := flag.String("config", "configuration.toml", "path to configuration file")
	migrate := flag.Bool("migrate", false, "run auto migrate? WARNING WILL DROP ALL TABLES")
	seed := flag.Bool("seed", false, "run seed? Create one hero for you to play it.")
	flag.Parse()

	app := setup(*configPath)

	if *migrate {
		if err := app.storage.AutoMigrateAll(); err != nil {
			app.log.Fatal(err)
		}
	}
	if *seed {
		if err := app.storage.Seed(); err != nil {
			app.log.Fatal(err)
		}
	}

	router := mux.NewRouter()

	superRouter := router.PathPrefix("/supers").Headers("Content-Type", "application/json").Subrouter()
	superRouter.HandleFunc("/", app.NewSuper).Methods(http.MethodPost)
	superRouter.HandleFunc("/", app.GetAll).Methods(http.MethodGet)
	superRouter.HandleFunc("/good", app.GetAllGood).Methods(http.MethodGet)
	superRouter.HandleFunc("/bad", app.GetAllBad).Methods(http.MethodGet)
	superRouter.HandleFunc("/search/{name}", app.GetByName).Methods(http.MethodGet)
	superRouter.HandleFunc("/{id}", app.GetByID).Methods(http.MethodGet)
	superRouter.HandleFunc("/{id}", app.DeleteByID).Methods(http.MethodDelete)

	superRouter.Use(alwaysJson)

	app.log.Println("Ready")

	app.log.Fatal(http.ListenAndServe(app.config.Server.Bind, router))
}

func setup(configPath string) *App {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	config, err := LoadConfig(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	storage, err := sql.New(config.Database.DSN)
	if err != nil {
		logger.Fatal(err)
	}

	return &App{storage: storage, log: logger, config: config}
}
