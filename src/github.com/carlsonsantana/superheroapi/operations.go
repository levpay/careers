package superheroapi

import "strings"

func ConvertSuperFromDatabase(rawSuper map[string]interface{}) *Super {
	return &Super{
		rawSuper["uuid"].(string),
		int(rawSuper["superheroapi_id"].(int64)),
		rawSuper["name"].(string),
		rawSuper["full_name"].(string),
		int(rawSuper["intelligence"].(int64)),
		int(rawSuper["power"].(int64)),
		rawSuper["occupation"].(string),
		rawSuper["image"].(string),
		strings.Split(rawSuper["groups"].(string), "|"),
		rawSuper["category"].(string),
		int(rawSuper["number_relatives"].(int64)),
	}
}

func AddSupersDatabase(supers []Super) error {
	db := GetDatabaseConnection()
	inserts := []map[string]interface{}{}
	for _, super := range supers {
		insert := map[string]interface{}{
			"uuid":             super.UUID,
			"superheroapi_id":  super.SuperHeroAPIID,
			"name":             super.Name,
			"full_name":        super.FullName,
			"intelligence":     super.Intelligence,
			"power":            super.Power,
			"occupation":       super.Occupation,
			"image":            super.Image,
			"groups":           strings.Join(super.Groups[:], "|"),
			"category":         super.Category,
			"number_relatives": super.NumberRelatives,
		}
		inserts = append(inserts, insert)
	}
	return db.Table("super").InsertBatch(inserts)
}

func GetSuperBySuperHeroAPIIDDatabase(superHeroAPIID int) *Super {
	db := GetDatabaseConnection()
	rawSuper, err := db.Table("super").Where("superheroapi_id", "=", superHeroAPIID).Get()
	if err != nil || len(rawSuper) == 0 {
		return nil
	}
	return ConvertSuperFromDatabase(rawSuper[0])
}

func ListSupersDatabase() []Super {
	db := GetDatabaseConnection()
	supers := []Super{}
	rawSupers, err := db.Table("super").Get()
	if err == nil {
		for _, rawSuper := range rawSupers {
			supers = append(supers, *ConvertSuperFromDatabase(rawSuper))
		}
	}
	return supers
}
