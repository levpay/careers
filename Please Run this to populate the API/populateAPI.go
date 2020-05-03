package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Baseado em : https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/

type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type Biography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type Appearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	Eyecolor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}
type Work struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}
type Connections struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}
type Image struct {
	Url string `json:"url"`
}
type SuperHeroData struct {
	Response    string      `json:"response"`
	Uuid        uint32      `json:"id"`
	Name        string      `json:"name""`
	Powerstats  Powerstats  `json:"powerstats""`
	Biography   Biography   `json:"biography"`
	Appearance  Appearance  `json:"appearance"`
	Work        Work        `json:"work"`
	Connections Connections `json:"connections"`
	Image       Image       `json:"image"`
}
type SuperOrVilan struct {
	Name             string `json:"name"             binding:"required" `
	FullName         string `json:"fullname"         binding:"required" `
	Alignment        string `json:"alignment"        binding:"required"`
	Intelligence     string `json:"intelligence"     binding:"required" `
	Power            string `json:"power"            binding:"required" `
	Occupation       string `json:"occupation"       binding:"required" `
	Image            string `json:"image"            binding:"required" `
	GroupAffiliation string `json:"groupAffiliation" binding:"required" `
	Relatives        string `json:"relatives"        binding:"required" `
}

func main() {
	fmt.Println("Starting the application...")
	 for i:=1;i<10;i++ {
	 	response, err := http.Get("https://superheroapi.com/api/" + os.Getenv("SUPERHERO_ID") + "/" + strconv.Itoa(i))
	 	if response.Status=="400"{
	 		break
		}
		 if err != nil {
			 fmt.Printf("The HTTP request failed with error %s\n", err)
			 break
			 return
		 } else {
			 var super SuperHeroData
			 data, _ := ioutil.ReadAll(response.Body)
			 json.Unmarshal([]byte(data), &super)
			 	superMyApi := SuperOrVilan{
			 		Name:             super.Name,
			 		FullName:         super.Biography.FullName,
			 		Alignment:        super.Biography.Alignment,
			 		Intelligence:     super.Powerstats.Intelligence,
			 		Power:            super.Powerstats.Power,
			 		Occupation:       super.Work.Occupation,
			 		Image:            super.Image.Url,
			 		GroupAffiliation: super.Connections.GroupAffiliation,
			 		Relatives:        super.Connections.Relatives,
			 	}
			 	if jsonValue, err := json.Marshal(superMyApi); err != nil {
			 		fmt.Println(err.Error())
			 	}else {
					response, err = http.Post("http://localhost:8080/create/user", "application/json", bytes.NewBuffer(jsonValue))
					if err != nil {
						fmt.Printf("The HTTP request failed with error %s\n", err)
					} else {
						data, _ := ioutil.ReadAll(response.Body)
						fmt.Println(string(data))
					}
				}

		 }
	 }

	fmt.Println("Terminating the application...")
}
