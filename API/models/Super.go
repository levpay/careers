package models

import "time"

type SuperOrVilan struct {
	Uuid             uint32   `gorm:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	Name             string   `json:"name" gorm:"unique;not null"`
	FullName         string   `json:"fullname"`
	Intelligence     string   `json:"intelligence"`
	Power            string   `json:"power"`
	Occupation       string   `json:"occupation"`
	Image            string   `json:"image"`
	GroupAffiliation []string `json:"groupAffiliation"`
	Relatives        []string `json:"relatives"`
}
type SuperOrVilanSearch struct {
	Uuid             uint32   `gorm:"primary_key"`
	Name             string   `json:"name" gorm:"unique;not null"`
	FullName         string   `json:"fullname"`
	Intelligence     string   `json:"intelligence"`
	Power            string   `json:"power"`
	Occupation       string   `json:"occupation"`
	Image            string   `json:"image"`
	GroupAffiliation []string `json:"groupAffiliation"`
	Relatives        int `json:"relatives"`
}
//Biding inputs
type SuperOrVilanInput struct {
	Name             string   `json:"name"             binding:"required" `
	FullName         string   `json:"fullname"         binding:"required" `
	Intelligence     string   `json:"intelligence"     binding:"required" `
	Power            string   `json:"power"            binding:"required" `
	Occupation       string   `json:"occupation"       binding:"required" `
	Image            string   `json:"image"            binding:"required" `
	GroupAffiliation []string `json:"groupAffiliation" binding:"required" `
	Relatives        []string `json:"relatives"        binding:"required" `
}
