package model

import (
	"github.com/google/uuid"
)

// Group contains name of group affiliation
type Group struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Name string    `gorm:"unique;not null"`

	Supers []*Super `gorm:"many2many:group_super;" json:"-"`
}
