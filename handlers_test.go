package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/dvdscripter/superheroapi/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// storageFake implements storage for handler testing
type storageFake struct {
	mustErr bool
}

var testSupers = []model.Super{
	model.Super{ID: uuid.Nil, Name: "Batman"},
	model.Super{ID: uuid.Nil, Name: "Superman"},
}

var errForTest = errors.New("erroring")

func (f *storageFake) CreateSuper(model.Super) (model.Super, error) {
	if f.mustErr {
		return model.Super{}, errForTest
	}
	return testSupers[0], nil
}

func (f *storageFake) ListAllSuper() ([]model.Super, error) {
	if f.mustErr {
		return nil, errForTest
	}
	return testSupers, nil
}

func (f *storageFake) ListAllGood() ([]model.Super, error) {
	if f.mustErr {
		return nil, errForTest
	}
	return testSupers, nil
}

func (f *storageFake) ListAllBad() ([]model.Super, error) {
	if f.mustErr {
		return nil, errForTest
	}
	return testSupers, nil
}

func (f *storageFake) FindByName(name string) (model.Super, error) {
	if f.mustErr {
		return model.Super{}, errForTest
	}
	return testSupers[0], nil
}

func (f *storageFake) FindByID(id string) (model.Super, error) {
	if f.mustErr {
		return model.Super{}, errForTest
	}
	return testSupers[0], nil
}
func (f *storageFake) DeleteByID(id string) error {
	if f.mustErr {
		return errForTest
	}
	return nil
}

func (f *storageFake) AutoMigrateAll() error {
	return nil
}

func (f *storageFake) Seed() error {
	return nil
}

func TestApp_NewSuper(t *testing.T) {
	type args struct {
		json io.Reader
	}
	tests := []struct {
		name string
		app  *App
		args args
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(`{"Name":"Batman"}`)},
			want: `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(`{"Name":"Batman"}`)},
			want: `{"error":"Oops something went wrong"}` + "\n",
		},
		{
			name: "error, no input",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(``)},
			want: `{"error":"Oops something went wrong"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", tt.args.json)
			w := httptest.NewRecorder()

			tt.app.NewSuper(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.NewSuper() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_GetAll(t *testing.T) {
	tests := []struct {
		name string
		app  *App
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""},{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}]` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `{"error":"Oops something went wrong"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			tt.app.GetAll(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.GetAll() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_GetAllGood(t *testing.T) {
	tests := []struct {
		name string
		app  *App
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""},{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}]` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `{"error":"Oops something went wrong"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			tt.app.GetAllGood(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.GetAllGood() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_GetAllBad(t *testing.T) {
	tests := []struct {
		name string
		app  *App
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""},{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}]` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: `{"error":"Oops something went wrong"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			tt.app.GetAllBad(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.GetAllBad() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_GetByName(t *testing.T) {
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"name": "Batman"}},
			want: `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: nil},
			want: `{"error":"var not found"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			req = mux.SetURLVars(req, tt.args.vars)

			tt.app.GetByName(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.GetByName() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_GetByID(t *testing.T) {
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "00000000-0000-0000-0000-000000000000"}},
			want: `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Morality":"","Alignment":""}` + "\n",
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: nil},
			want: `{"error":"var not found"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			req = mux.SetURLVars(req, tt.args.vars)

			tt.app.GetByID(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.GetByID() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}

func TestApp_DeleteByID(t *testing.T) {
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want string
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "00000000-0000-0000-0000-000000000000"}},
			want: ``,
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: nil},
			want: `{"error":"var not found"}` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			req = mux.SetURLVars(req, tt.args.vars)

			tt.app.DeleteByID(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("app.DeleteByID() = '%s', want '%v'", got, tt.want)
			}

		})
	}
}
