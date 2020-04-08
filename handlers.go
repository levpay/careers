package main

import (
	"net/http"

	"github.com/dvdscripter/careers/model"
)

var genericErr = map[string]string{"error": "Oops something went wrong"}

func (app *App) NewSuper(w http.ResponseWriter, r *http.Request) {
	var super model.Super
	if err := fromJSON(r.Body, &super); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	super, err := app.storage.CreateSuper(super)
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := toJSON(w, &super); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) GetAll(w http.ResponseWriter, r *http.Request) {
	supers, err := app.storage.ListAllSuper()
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) GetAllGood(w http.ResponseWriter, r *http.Request) {
	supers, err := app.storage.ListAllSuper()
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) GetAllBad(w http.ResponseWriter, r *http.Request) {
	supers, err := app.storage.ListAllBad()
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) GetByName(w http.ResponseWriter, r *http.Request) {
	name, err := getVar(r, "name")
	if err != nil {
		missing := genericErr
		missing["error"] = ErrVarNotFound.Error()
		toJSON(w, missing)
		app.log.Println(err)
		return
	}

	supers, err := app.storage.FindByName(name)
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getVar(r, "id")
	if err != nil {
		missing := genericErr
		missing["error"] = ErrVarNotFound.Error()
		toJSON(w, missing)
		app.log.Println(err)
		return
	}

	supers, err := app.storage.FindByID(id)
	if err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}

func (app *App) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, err := getVar(r, "id")
	if err != nil {
		missing := genericErr
		missing["error"] = ErrVarNotFound.Error()
		toJSON(w, missing)
		app.log.Println(err)
		return
	}

	if err := app.storage.DeleteByID(id); err != nil {
		toJSON(w, genericErr)
		app.log.Println(err)
		return
	}
}
