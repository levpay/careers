package main

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "working",
			args: args{"testdata/configuration.toml"},
			want: &Config{
				Server:   Server{Bind: "0.0.0.0:8000"},
				Database: Database{DSN: "postgres://@localhost/super"},
			},
			wantErr: false,
		},
		{
			name:    "file not found, I hope...",
			args:    args{"testdata/configuration6116541.toml"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
