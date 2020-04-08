package model

import (
	"github.com/google/uuid"
)

// Super contains villain or hero information
type Super struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name         string    `gorm:"unique;not null"`
	FullName     string
	Intelligence int
	Power        int
	Occupation   string
	Image        string
	Parents      int    `gorm:"-"`
	Morality     string `gorm:"-"`

	Alignment string
	Groups    []*Group    `gorm:"many2many:group_super;" json:",omitempty"`
	Relatives []*Relative `gorm:"many2many:relative_super;" json:",omitempty"`
}

func (model *Super) AfterFind() error {
	if model.Relatives != nil {
		model.Parents = len(model.Relatives)
	}

	return nil
}
