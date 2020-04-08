package storage

import "github.com/dvdscripter/superheroapi/model"

// SuperStorage is our contract which must be meet to interact with persistence.
// All relationships should be preloaded before returning
type SuperStorage interface {
	// CreateSuper should create and handle model.Super relationships creation
	CreateSuper(model.Super) (model.Super, error)
	// ListAllSuper should return all Supers
	ListAllSuper() ([]model.Super, error)
	// ListAllGood should return all Supers with alignment == "good"
	ListAllGood() ([]model.Super, error)
	// ListAllBad should return all Supers with alignment == "bad"
	ListAllBad() ([]model.Super, error)
	// FindByName should return the first record filtered by name
	FindByName(string) (model.Super, error)
	// FindByID should return the first record filtered by id
	FindByID(string) (model.Super, error)
	// DeleteByID should remove Super and their references without clearing
	DeleteByID(string) error

	// AutoMigrateAll should perform migrations based on model.* definitions
	AutoMigrateAll() error
	// Seed Create initial data
	Seed() error
}
