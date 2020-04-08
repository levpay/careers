package sql

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dvdscripter/superheroapi/model"
	"github.com/google/uuid"
)

func TestDB_CreateSuper(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	super := model.Super{Name: "Batman"}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "supers" ("name","full_name","intelligence","power","occupation","image","alignment") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "supers"."id"`)).
		WithArgs(super.Name, super.FullName, super.Intelligence, super.Power, super.Occupation, super.Image, super.Alignment).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.Nil))
	mock.ExpectCommit()

	if _, err := storage.CreateSuper(super); err != nil {
		t.Errorf("DB.CreateSuper() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.CreateSuper() error = %v", err)
		return
	}

}

func TestDB_ListAllSuper(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.Nil, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers"`)).WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	if _, err := storage.ListAllSuper(); err != nil {
		t.Errorf("DB.ListAllSuper() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.ListAllSuper() error = %v", err)
		return
	}
}

func TestDB_ListAllGood(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.Nil, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers" WHERE (alignment = $1)`)).WithArgs("good").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	if _, err := storage.ListAllGood(); err != nil {
		t.Errorf("DB.ListAllGood() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.ListAllGood() error = %v", err)
		return
	}
}

func TestDB_ListAllBad(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.Nil, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers" WHERE (alignment = $1)`)).WithArgs("bad").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	if _, err := storage.ListAllBad(); err != nil {
		t.Errorf("DB.ListAllBad() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.ListAllBad() error = %v", err)
		return
	}
}

func TestDB_FindByName(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.Nil, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers" WHERE (name = $1) LIMIT 1`)).WithArgs("batman").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	if _, err := storage.FindByName("batman"); err != nil {
		t.Errorf("DB.FindByName() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.FindByName() error = %v", err)
		return
	}
}

func TestDB_FindByID(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.Nil, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers" WHERE (id = $1)`)).WithArgs(uuid.Nil).WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	if _, err := storage.FindByID("00000000-0000-0000-0000-000000000000"); err != nil {
		t.Errorf("DB.FindByID() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.FindByID() error = %v", err)
		return
	}
}

func TestDB_DeleteByID(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "image"}).
		AddRow(uuid.NameSpaceURL, "batman", "bruce weiner", 10, 10, "hero", "")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "supers" WHERE (id = $1) LIMIT 1`)).WithArgs(uuid.NameSpaceURL).WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"groups\".+").WillReturnRows(rows)
	mock.ExpectQuery("SELECT \\* FROM \"relatives\".+").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"relative_super\".+").WithArgs(uuid.NameSpaceURL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"group_super\".+").WithArgs(uuid.NameSpaceURL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "supers" WHERE "supers"."id" = $1`)).WithArgs(uuid.NameSpaceURL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := storage.DeleteByID(uuid.NameSpaceURL.String()); err != nil {
		t.Errorf("DB.DeleteByID() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.DeleteByID() error = %v", err)
		return
	}
}

func TestDB_findOrCreateGroupFind(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	group := model.Group{Name: "League of Justice"}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(uuid.NameSpaceURL, group.Name)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups" WHERE (name = $1) ORDER BY "groups"."id" ASC LIMIT 1`)).WithArgs(group.Name).WillReturnRows(rows)

	if _, err := storage.findOrCreateGroup(group); err != nil {
		t.Errorf("DB.findOrCreateGroup() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.findOrCreateGroup() error = %v", err)
		return
	}
}

func TestDB_findOrCreateGroupCreate(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	group := model.Group{Name: "League of Justice"}

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups" WHERE (name = $1) ORDER BY "groups"."id" ASC LIMIT 1`)).WithArgs(group.Name).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "groups" ("name") VALUES ($1) RETURNING "groups"."id"`)).
		WithArgs(group.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.Nil))
	mock.ExpectCommit()

	if _, err := storage.findOrCreateGroup(group); err != nil {
		t.Errorf("DB.findOrCreateGroup() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.findOrCreateGroup() error = %v", err)
		return
	}
}

func TestDB_findOrCreateRelativeFind(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	relative := model.Relative{
		Name:    "Alfred Pennyworth",
		Kinship: "former guardian",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "kinship"}).
		AddRow(uuid.NameSpaceURL, relative.Name, relative.Kinship)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "relatives" WHERE (name = $1) ORDER BY "relatives"."id" ASC LIMIT 1`)).
		WithArgs(relative.Name).WillReturnRows(rows)

	if _, err := storage.findOrCreateRelative(relative); err != nil {
		t.Errorf("DB.findOrCreateRelative() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.findOrCreateRelative() error = %v", err)
		return
	}
}

func TestDB_findOrCreateRelativeCreate(t *testing.T) {
	t.Parallel()
	storage, mock := newTestDB(t)
	defer storage.DB.Close()

	relative := model.Relative{
		Name:    "Alfred Pennyworth",
		Kinship: "former guardian",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "kinship"})

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "relatives" WHERE (name = $1) ORDER BY "relatives"."id" ASC LIMIT 1`)).WithArgs(relative.Name).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "relatives" ("name","kinship") VALUES ($1,$2) RETURNING "relatives"."id"`)).
		WithArgs(relative.Name, relative.Kinship).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.Nil))
	mock.ExpectCommit()

	if _, err := storage.findOrCreateRelative(relative); err != nil {
		t.Errorf("DB.findOrCreateRelative() error = %v", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("DB.findOrCreateRelative() error = %v", err)
		return
	}
}
