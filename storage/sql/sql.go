package sql

import (
	"github.com/dvdscripter/superheroapi/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func New(dbpath string) (*DB, error) {
	db, err := gorm.Open("postgres", dbpath)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) AutoMigrateAll() error {
	if err := db.DropTableIfExists(&model.Super{}, &model.Group{}, &model.Relative{}, "relative_super", "group_super").Error; err != nil {
		return err
	}
	return db.AutoMigrate(&model.Super{}, &model.Group{}, &model.Relative{}).Error
}

func (db *DB) Close() error {
	return db.DB.Close()
}
