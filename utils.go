package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var ErrVarNotFound = errors.New("var not found")

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
		return "", ErrVarNotFound
	}
	return named, nil
}
