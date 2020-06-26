package superheroapi

import "strings"

func AddSupersDatabase(supers []Super) error {
	db := GetDatabaseConnection()
	inserts := []map[string]interface{}{}
	for _, super := range supers {
		insert := map[string]interface{}{
			"uuid":             super.UUID,
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
