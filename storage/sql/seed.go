package sql

import (
	"github.com/dvdscripter/careers/model"
)

func (db *DB) Seed() error {

	batman := model.Super{
		Name:         "Batman",
		FullName:     "Bruce Wayne",
		Intelligence: 100,
		Power:        47,
		Occupation:   "Businessman",
		Image:        "httpss://www.superherodb.com/pictures2/portraits/10/100/639.jpg",
		Alignment:    "good",
		Groups: []*model.Group{
			{Name: "Batman Family"},
			{Name: "Batman Incorporated"},
			{Name: "Justice League"},
			{Name: "Outsiders"},
			{Name: "Wayne Enterprises"},
			{Name: "Club of Heroes"},
			{Name: "formerly White Lantern Corps"},
			{Name: "Sinestro Corps"},
		},
		Relatives: []*model.Relative{
			{
				Name:    "Damian Wayne",
				Kinship: "son",
			},
			{
				Name:    "Dick Grayson",
				Kinship: "adopted son",
			},
			{
				Name:    "Tim Drake",
				Kinship: "adopted son",
			},
			{
				Name:    "Jason Todd",
				Kinship: "adopted son",
			},
			{
				Name:    "Cassandra Cain",
				Kinship: "adopted ward",
			},
			{
				Name:    "Martha Wayne",
				Kinship: "mother",
			},
			{
				Name:    "Thomas Wayne",
				Kinship: "father",
			},
			{
				Name:    "Alfred Pennyworth",
				Kinship: "former guardian",
			},
			{
				Name:    "Roderick Kane",
				Kinship: "grandfather",
			},
			{
				Name:    "Elizabeth Kane",
				Kinship: "grandmother",
			},
			{
				Name:    "Nathan Kane",
				Kinship: "uncle",
			},
			{
				Name:    "Simon Hurt",
				Kinship: "ancestor",
			},
		},
	}

	if _, err := db.CreateSuper(batman); err != nil {
		return err
	}

	return nil
}
