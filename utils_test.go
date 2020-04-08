package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func Test_alwaysJson(t *testing.T) {

	want := "application/json"

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		toJSON(w, map[string]string{"error": "not found"})
	}

	alwaysJson(http.HandlerFunc(handler)).ServeHTTP(w, req)

	resp := w.Result()
	if got := resp.Header.Get("Content-Type"); got != want {
		t.Errorf("alwaysJson() %v, want %v", got, want)
		return
	}

}

func Test_getVar(t *testing.T) {

	type args struct {
		vars map[string]string
		key  string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "working",
			args: args{
				vars: map[string]string{"name": "batman"},
				key:  "name",
			},
			want:    "batman",
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				vars: map[string]string{"name": "superman"},
				key:  "id",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
				return
			}
			req = mux.SetURLVars(req, tt.args.vars)

			got, err := getVar(req, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVar() = %v, want %v", got, tt.want)
			}

		})
	}

}
