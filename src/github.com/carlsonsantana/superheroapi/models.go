package superheroapi

import (
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

type Super struct {
	UUID            string `json:"uuid"`
	SuperHeroAPIID  int
	Name            string   `json:"name"`
	FullName        string   `json:"full-name"`
	Intelligence    int      `json:"intelligence"`
	Power           int      `json:"power"`
	Occupation      string   `json:"occupation"`
	Image           string   `json:"image"`
	Groups          []string `json:"groups"`
	Category        string   `json:"category"`
	NumberRelatives int      `json:"number-relatives"`
}

type SuperHeroAPIPowerStatus struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type SuperHeroAPIBiography struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceBirth      string   `json:"place-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type SuperHeroAPIAppearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eye-color"`
	HairColor string   `json:"hair-color"`
}

type SuperHeroAPIWork struct {
	Occupation string `json:"occupation"`
	Case       string `json:"case"`
}

type SuperHeroAPIConnections struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}

type SuperHeroAPIImage struct {
	URL string `json:"url"`
}

type SuperHeroAPISuper struct {
	ID          string                  `json:"id"`
	Name        string                  `json:"name"`
	Powerstats  SuperHeroAPIPowerStatus `json:"powerstats"`
	Biography   SuperHeroAPIBiography   `json:"biography"`
	Appearance  SuperHeroAPIAppearance  `json:"appearance"`
	Work        SuperHeroAPIWork        `json:"work"`
	Connections SuperHeroAPIConnections `json:"connections"`
	Image       SuperHeroAPIImage       `json:"image"`
}

func ConvertSuperHeroAPIResponseToSuper(
	superHeroAPIResponse *SuperHeroAPIResponse,
) []Super {
	if superHeroAPIResponse.Error != "" {
		return nil
	}
	supersHeroAPIResults := superHeroAPIResponse.Results
	supers := []Super{}
	for _, superHeroAPIResult := range supersHeroAPIResults {
		superHeroAPIID, _ := strconv.Atoi(superHeroAPIResult.ID)
		super := GetSuperBySuperHeroAPIIDDatabase(superHeroAPIID)
		if super == nil {
			intelligence, _ := strconv.Atoi(
				superHeroAPIResult.Powerstats.Intelligence,
			)
			power, _ := strconv.Atoi(superHeroAPIResult.Powerstats.Power)
			var category string
			if superHeroAPIResult.Biography.Alignment == "good" {
				category = "hero"
			} else if superHeroAPIResult.Biography.Alignment == "bad" {
				category = "villain"
			} else {
				category = "neutral"
			}
			super = &Super{
				uuid.NewV4().String(),
				superHeroAPIID,
				superHeroAPIResult.Name,
				superHeroAPIResult.Biography.FullName,
				intelligence,
				power,
				superHeroAPIResult.Work.Occupation,
				superHeroAPIResult.Image.URL,
				strings.Split(
					superHeroAPIResult.Connections.GroupAffiliation,
					", ",
				),
				category,
				len(strings.Split(superHeroAPIResult.Connections.Relatives, ", ")),
			}
		}
		supers = append(supers, *super)
	}
	return supers
}
