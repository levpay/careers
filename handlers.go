package main

import (
	"net/http"

	"github.com/dvdscripter/careers/model"
)

var genericErr = "Oops something went wrong"

func (app *App) NewSuper(w http.ResponseWriter, r *http.Request) {
	var super model.Super
	if err := fromJSON(r.Body, &super); err != nil {
		app.Error(w, err, genericErr, http.StatusBadRequest)
		return
	}

	super, err := app.storage.CreateSuper(super)
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	// Location with non 3xx http status code does not seems well supported
	// curl will ignore this following even with -L flag
	w.Header().Add("Location", "/supers/"+super.ID.String())
	w.WriteHeader(http.StatusCreated)
}

func (app *App) GetAll(w http.ResponseWriter, r *http.Request) {

	if alignment := r.URL.Query().Get("alignment"); alignment != "" {
		switch alignment {
		case "good":
			app.GetAllGood(w, r)
			return
		case "bad":
			app.GetAllBad(w, r)
			return
		}
	}
	if name := r.URL.Query().Get("name"); name != "" {
		app.GetByName(w, r, name)
		return
	}

	supers, err := app.storage.ListAllSuper()
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}
}

func (app *App) GetAllGood(w http.ResponseWriter, r *http.Request) {
	supers, err := app.storage.ListAllGood()
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}
}

func (app *App) GetAllBad(w http.ResponseWriter, r *http.Request) {
	supers, err := app.storage.ListAllBad()
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}
}

func (app *App) GetByName(w http.ResponseWriter, r *http.Request, name string) {

	supers, err := app.storage.FindByName(name)
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}
}

func (app *App) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getVar(r, "id")
	if err != nil {
		app.Error(w, err, "missing id parameter", http.StatusBadRequest)
		return
	}

	supers, err := app.storage.FindByID(id)
	if err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	if err := toJSON(w, &supers); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}
}

func (app *App) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id, err := getVar(r, "id")
	if err != nil {
		app.Error(w, err, "missing id parameter", http.StatusBadRequest)
		return
	}

	// delete should really by idempotent? trying to remove non existing content
	// should return 404? API must be informative, so I take this way
	if err := app.storage.DeleteByID(id); err != nil {
		app.Error(w, err, genericErr, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
