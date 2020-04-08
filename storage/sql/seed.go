package sql

import (
	"github.com/dvdscripter/superheroapi/model"
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
			&model.Group{Name: "Batman Family"},
			&model.Group{Name: "Batman Incorporated"},
			&model.Group{Name: "Justice League"},
			&model.Group{Name: "Outsiders"},
			&model.Group{Name: "Wayne Enterprises"},
			&model.Group{Name: "Club of Heroes"},
			&model.Group{Name: "formerly White Lantern Corps"},
			&model.Group{Name: "Sinestro Corps"},
		},
		Relatives: []*model.Relative{
			&model.Relative{
				Name:    "Damian Wayne",
				Kinship: "son",
			},
			&model.Relative{
				Name:    "Dick Grayson",
				Kinship: "adopted son",
			},
			&model.Relative{
				Name:    "Tim Drake",
				Kinship: "adopted son",
			},
			&model.Relative{
				Name:    "Jason Todd",
				Kinship: "adopted son",
			},
			&model.Relative{
				Name:    "Cassandra Cain",
				Kinship: "adopted ward",
			},
			&model.Relative{
				Name:    "Martha Wayne",
				Kinship: "mother",
			},
			&model.Relative{
				Name:    "Thomas Wayne",
				Kinship: "father",
			},
			&model.Relative{
				Name:    "Alfred Pennyworth",
				Kinship: "former guardian",
			},
			&model.Relative{
				Name:    "Roderick Kane",
				Kinship: "grandfather",
			},
			&model.Relative{
				Name:    "Elizabeth Kane",
				Kinship: "grandmother",
			},
			&model.Relative{
				Name:    "Nathan Kane",
				Kinship: "uncle",
			},
			&model.Relative{
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
