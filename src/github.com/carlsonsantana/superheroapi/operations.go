package superheroapi

import "strings"

func ConvertSuperFromDatabase(rawSuper []map[string]interface{}) *Super {
	if len(rawSuper) == 0 {
		return nil
	}
	return &Super{
		rawSuper[0]["uuid"].(string),
		int(rawSuper[0]["superheroapi_id"].(int64)),
		rawSuper[0]["name"].(string),
		rawSuper[0]["full_name"].(string),
		int(rawSuper[0]["intelligence"].(int64)),
		int(rawSuper[0]["power"].(int64)),
		rawSuper[0]["occupation"].(string),
		rawSuper[0]["image"].(string),
		strings.Split(rawSuper[0]["groups"].(string), "|"),
		rawSuper[0]["category"].(string),
		int(rawSuper[0]["number_relatives"].(int64)),
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
	if err != nil {
		return nil
	}
	return ConvertSuperFromDatabase(rawSuper)
}
