package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (app *App) Error(w http.ResponseWriter, err error, publicMessage string, statusCode int) {

	switch errors.Cause(err) {
	case gorm.ErrRecordNotFound:
		w.WriteHeader(http.StatusNotFound)
		toJSON(w, map[string]string{"error": "record not found"})
	default:
		w.WriteHeader(statusCode)
		toJSON(w, map[string]string{"error": publicMessage})
	}

	app.log.Printf("%+v\n", err)
}

func alwaysJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func toJSON(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(&data)
}

func fromJSON(r io.Reader, dest interface{}) error {
	return json.NewDecoder(r).Decode(&dest)
}

func getVar(r *http.Request, key string) (string, error) {
	vars := mux.Vars(r)
	named, exist := vars[key]
	if !exist {
		return "", errors.New("var not found")
	}
	return named, nil
}
