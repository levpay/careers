package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/dvdscripter/careers/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// storageFake implements storage for handler testing
type storageFake struct {
	mustErr bool
	err     error
}

var testSupers = []model.Super{
	{ID: uuid.Nil, Name: "Batman", Alignment: "good"},
	{ID: uuid.Nil, Name: "Superman", Alignment: "bad"},
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
	return []model.Super{testSupers[0]}, nil
}

func (f *storageFake) ListAllBad() ([]model.Super, error) {
	if f.mustErr {
		return nil, errForTest
	}
	return []model.Super{testSupers[1]}, nil
}

func (f *storageFake) FindByName(name string) (model.Super, error) {
	if f.mustErr {
		if f.err == nil {
			return model.Super{}, errForTest
		}
		return model.Super{}, f.err
	}
	return testSupers[0], nil
}

func (f *storageFake) FindByID(id string) (model.Super, error) {
	if f.mustErr {
		if f.err == nil {
			return model.Super{}, errForTest
		}
		return model.Super{}, f.err
	}
	return testSupers[0], nil
}
func (f *storageFake) DeleteByID(id string) error {
	if f.mustErr {
		if f.err == nil {
			return errForTest
		}
		return f.err
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
	type want struct {
		content    string
		statusCode int
		location   string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(`{"Name":"Batman"}`)},
			want: want{
				statusCode: http.StatusCreated,
				location:   `/supers/00000000-0000-0000-0000-000000000000`,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(`{"Name":"Batman"}`)},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "error, no input",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{json: strings.NewReader(``)},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusBadRequest,
			},
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

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.NewSuper() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if location := resp.Header.Get("Location"); location != tt.want.location {
				t.Errorf("app.NewSuper() = %v, want %v", location, tt.want.location)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.NewSuper() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_GetAll(t *testing.T) {
	type args struct {
		alignment string
		name      string
	}
	type want struct {
		content    string
		statusCode int
	}
	tests := []struct {
		name string
		app  *App
		args args
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"},{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"bad"}]` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "working good alignment",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{alignment: "good"},
			want: want{
				content:    `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"}]` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "working bad alignment",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{alignment: "bad"},
			want: want{
				content:    `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"bad"}]` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "working name filter",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{name: "Batman"},
			want: want{
				content:    `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"}` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "error, name not found",
			app: &App{
				storage: &storageFake{mustErr: true, err: gorm.ErrRecordNotFound},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{name: "david"},
			want: want{
				content:    `{"error":"record not found"}` + "\n",
				statusCode: http.StatusNotFound,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			params := url.Values{}
			params.Set("alignment", tt.args.alignment)
			params.Set("name", tt.args.name)
			req.URL.RawQuery = params.Encode()

			w := httptest.NewRecorder()

			tt.app.GetAll(w, req)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.GetAll() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.GetAll() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_GetAllGood(t *testing.T) {
	type want struct {
		content    string
		statusCode int
	}
	tests := []struct {
		name string
		app  *App
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"}]` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusInternalServerError,
			},
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

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.GetAllGood() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.GetAllGood() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_GetAllBad(t *testing.T) {
	type want struct {
		content    string
		statusCode int
	}
	tests := []struct {
		name string
		app  *App
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `[{"ID":"00000000-0000-0000-0000-000000000000","Name":"Superman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"bad"}]` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusInternalServerError,
			},
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

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.GetAllBad() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.GetAllBad() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_GetByName(t *testing.T) {
	type args struct {
		name string
	}
	type want struct {
		content    string
		statusCode int
	}
	tests := []struct {
		name string
		app  *App
		args args
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{name: "Batman"},
			want: want{
				content:    `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"}` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "error generic",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{name: "lex"},
			want: want{
				content:    `{"error":"Oops something went wrong"}` + "\n",
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "error not found",
			app: &App{
				storage: &storageFake{mustErr: true, err: gorm.ErrRecordNotFound},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{name: "lex"},
			want: want{
				content:    `{"error":"record not found"}` + "\n",
				statusCode: http.StatusNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			w := httptest.NewRecorder()

			tt.app.GetByName(w, req, tt.args.name)

			resp := w.Result()
			got, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.GetByName() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.GetByName() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_GetByID(t *testing.T) {
	type want struct {
		content    string
		statusCode int
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "00000000-0000-0000-0000-000000000000"}},
			want: want{
				content:    `{"ID":"00000000-0000-0000-0000-000000000000","Name":"Batman","FullName":"","Intelligence":0,"Power":0,"Occupation":"","Image":"","Parents":0,"Alignment":"good"}` + "\n",
				statusCode: http.StatusOK,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: nil},
			want: want{
				content:    `{"error":"missing id parameter"}` + "\n",
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error not found",
			app: &App{
				storage: &storageFake{mustErr: true, err: gorm.ErrRecordNotFound},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "10000000-0000-0000-0000-000000000001"}},
			want: want{
				content:    `{"error":"record not found"}` + "\n",
				statusCode: http.StatusNotFound,
			},
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

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.GetByID() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.GetByID() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}

func TestApp_DeleteByID(t *testing.T) {
	type want struct {
		content    string
		statusCode int
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name string
		app  *App
		args args
		want want
	}{
		{
			name: "working",
			app: &App{
				storage: &storageFake{mustErr: false},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "00000000-0000-0000-0000-000000000000"}},
			want: want{
				content:    ``,
				statusCode: http.StatusNoContent,
			},
		},
		{
			name: "error",
			app: &App{
				storage: &storageFake{mustErr: true},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: nil},
			want: want{
				content:    `{"error":"missing id parameter"}` + "\n",
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error, not found",
			app: &App{
				storage: &storageFake{mustErr: true, err: gorm.ErrRecordNotFound},
				log:     log.New(ioutil.Discard, "", log.LstdFlags),
			},
			args: args{vars: map[string]string{"id": "10000000-0000-0000-0000-000000000001"}},
			want: want{
				content:    `{"error":"record not found"}` + "\n",
				statusCode: http.StatusNotFound,
			},
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

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("app.DeleteByID() = %v, want %v", resp.StatusCode, tt.want.statusCode)
			}

			if !reflect.DeepEqual(string(got), tt.want.content) {
				t.Errorf("app.DeleteByID() = '%s', want '%v'", got, tt.want.content)
			}

		})
	}
}
