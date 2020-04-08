package main

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

type Database struct {
	DSN string
}

type Server struct {
	Bind string
}

// Config will contain any information which change between builds/deploy
// environments
type Config struct {
	Database Database
	Server   Server
}

// LoadConfig crate a new Config from path toml file
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	confContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(confContent, config); err != nil {
		return nil, err
	}

	return config, err
}
