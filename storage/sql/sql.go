package sql

import (
	"github.com/dvdscripter/careers/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type DB struct {
	*gorm.DB
}

func New(dbpath string) (*DB, error) {
	db, err := gorm.Open("postgres", dbpath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open db")
	}

	return &DB{db}, nil
}

func (db *DB) AutoMigrateAll() error {
	if err := db.DropTableIfExists(&model.Super{}, &model.Group{}, &model.Relative{}, "relative_super", "group_super").Error; err != nil {
		return errors.Wrap(err, "cannot drop tables for migration")
	}
	return errors.Wrap(db.AutoMigrate(&model.Super{}, &model.Group{}, &model.Relative{}).Error, "cannot auto migrate")
}

func (db *DB) Close() error {
	return errors.Wrap(db.DB.Close(), "cannot close db")
}
