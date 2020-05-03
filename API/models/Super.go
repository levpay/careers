package models

import "time"

type SuperOrVilan struct {
	Uuid             uint32   `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	Name             string   `json:"name" gorm:"unique"`
	FullName         string   `json:"fullname"`
	Intelligence     string   `json:"intelligence"`
	Alignment        string   `json:"alignment"`
	Power            string   `json:"power"`
	Occupation       string   `json:"occupation"`
	Image            string   `json:"image"`
	GroupAffiliation string   `json:"groupAffiliation" binding:"required"`
	Relatives        string   `json:"groupAffiliation" binding:"relatives"`
	Deleted          bool     `json:"deleted"`
}
type SuperOrVilanSearch struct {
	Uuid             uint32   `gorm:"primary_key"`
	Name             string   `json:"name" gorm:"unique;not null"`
	FullName         string   `json:"fullname"`
	Intelligence     string   `json:"intelligence"`
	Power            string   `json:"power"`
	Occupation       string   `json:"occupation"`
	Image            string   `json:"image"`
	GroupAffiliation string   `json:"groupAffiliation"`
	Relatives        int `json:"relatives"`
}
//Biding inputs
type SuperOrVilanInput struct {
	Name             string            `json:"name"              `
	FullName         string            `json:"fullname"          `
	Alignment        string            `json:"alignment"         `
	Intelligence     string            `json:"intelligence"      `
	Power            string            `json:"power"             `
	Occupation       string            `json:"occupation"        `
	Image            string            `json:"image"             `
	GroupAffiliation string            `json:"groupAffiliation"  `
	Relatives        string            `json:"relatives"         `
}
