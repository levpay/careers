package superheroapi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arthurkushman/buildsqlx"
)

var ExactStringParameters []string = []string{"uuid", "image", "category"}
var IlikeStringParameters []string = []string{
	"name",
	"full-name",
	"occupation",
	"groups",
}
var ExactNumberParameters []string = []string{"superheroapi-id"}
var ComparativeNumberParameters []string = []string{
	"intelligence",
	"power",
	"number-relatives",
}

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

func filterQueryByParameter(
	qDB *buildsqlx.DB,
	filterParameter string,
	filterValue string,
) *buildsqlx.DB {
	for _, parameterName := range ExactStringParameters {
		if parameterName == filterParameter {
			return qDB.AndWhere(
				/*
				 * Because buildsqlx bug I need to add a single quote before and
				 * after the column name.
				 */
				fmt.Sprintf(
					"'''' || %s || ''''",
					strings.ReplaceAll(filterParameter, "-", "_"),
				),
				"=",
				filterValue,
			)
		}
	}
	for _, parameterName := range IlikeStringParameters {
		if parameterName == filterParameter {
			return qDB.AndWhere(
				/*
				 * Because buildsqlx bug I need to add a single quote before and
				 * after the column name.
				 */
				fmt.Sprintf(
					"'''' || %s || ''''",
					strings.ReplaceAll(filterParameter, "-", "_"),
				),
				"ILIKE",
				filterValue,
			)
		}
	}
	for _, parameterName := range ExactNumberParameters {
		if parameterName == filterParameter {
			filterValueNumber, _ := strconv.Atoi(filterValue)
			return qDB.AndWhere(
				strings.ReplaceAll(filterParameter, "-", "_"),
				"=",
				filterValueNumber,
			)
		}
	}
	for _, parameterName := range ComparativeNumberParameters {
		if parameterName == filterParameter {
			operation := "="
			var filterValueNumber int
			if filterValue[0:1] == "<" {
				if filterValue[1:2] == "=" {
					operation = "<="
					filterValueNumber, _ = strconv.Atoi(filterValue[2:])
				} else {
					operation = "<"
					filterValueNumber, _ = strconv.Atoi(filterValue[1:])
				}
			} else if filterValue[0:1] == ">" {
				if filterValue[1:2] == "=" {
					operation = ">="
					filterValueNumber, _ = strconv.Atoi(filterValue[2:])
				} else {
					operation = ">"
					filterValueNumber, _ = strconv.Atoi(filterValue[1:])
				}
			} else {
				filterValueNumber, _ = strconv.Atoi(filterValue)
			}
			return qDB.AndWhere(
				strings.ReplaceAll(filterParameter, "-", "_"),
				operation,
				filterValueNumber,
			)
		}
	}
	return nil
}

func AddSupersDatabase(supers []Super) {
	db := GetDatabaseConnection()
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
		db.Table("super").Insert(insert)
	}
}

func GetSuperBySuperHeroAPIIDDatabase(superHeroAPIID int) *Super {
	db := GetDatabaseConnection()
	rawSuper, err := db.Table("super").Where(
		"superheroapi_id",
		"=",
		superHeroAPIID,
	).Get()
	if err != nil || len(rawSuper) == 0 {
		return nil
	}
	return ConvertSuperFromDatabase(rawSuper[0])
}

func GetSuperByUUID(uuid string) *Super {
	db := GetDatabaseConnection()
	rawSuper, err := db.Table("super").Where(
		/*
		 * Because buildsqlx bug I need to add a single quote before and
		 * after the column name.
		 */
		"'''' || uuid || ''''",
		"=",
		uuid,
	).Get()
	if err != nil || len(rawSuper) == 0 {
		return nil
	}
	return ConvertSuperFromDatabase(rawSuper[0])
}

func ListSupersDatabase(filters map[string]string) []Super {
	db := GetDatabaseConnection()
	supers := []Super{}
	qDB := db.Table("super").Where("1", "=", 1)

	for filterParameter, filterValue := range filters {
		qDB = filterQueryByParameter(qDB, filterParameter, filterValue)
	}

	rawSupers, err := qDB.Get()
	if err == nil {
		for _, rawSuper := range rawSupers {
			supers = append(supers, *ConvertSuperFromDatabase(rawSuper))
		}
	}
	return supers
}

func DeleteSuperDatabase(super *Super) {
	db := GetDatabaseConnection()
	db.Table("super").Where("uuid", "=", super.UUID).Delete()
}
