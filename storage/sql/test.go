package sql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func newTestDB(t *testing.T) (*DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	testDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatal(err)
	}

	return &DB{testDB}, mock

}
