package model

import (
	"github.com/google/uuid"
)

// Relative contains super relatives
type Relative struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	Name    string    `gorm:"not null"`
	Kinship string    `gorm:"not null"`

	Supers []*Super `gorm:"many2many:relative_super;" json:"-"`
}
