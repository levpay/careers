package superheroapi

import (
	"fmt"
	"net/http"
	"strconv"
)

type ValidationError struct {
	httpStatus int
	message    string
}

func ValidateParameterRequired(name string, value string) *ValidationError {
	if value == "" {
		return &ValidationError{
			http.StatusUnprocessableEntity,
			fmt.Sprintf("Parameter '%s' required.", name),
		}
	}
	return nil
}

func ValidateErrorInSuperHeroAPI(
	response *http.Response,
	err error,
) *ValidationError {
	if err != nil || response.StatusCode != http.StatusOK {
		return &ValidationError{http.StatusInternalServerError, "Erro interno"}
	}
	return nil
}

func ValidateSuperExistsInSuperHeroAPI(
	superHeroAPIResponse *SuperHeroAPIResponse,
	name string,
) *ValidationError {
	if superHeroAPIResponse.Error != "" {
		return &ValidationError{
			http.StatusFailedDependency,
			fmt.Sprintf(
				"Não foi possível encontrar um super com o nome '%s'.",
				name,
			),
		}
	}
	return nil
}

func ValidateInvalidFilterParameters(
	filters map[string]string,
) *ValidationError {
	allowedParameters := []string{
		"uuid",
		"superheroapi-id",
		"name",
		"full-name",
		"intelligence",
		"power",
		"occupation",
		"image",
		"groups",
		"category",
		"number-relatives",
	}
	var isAllowedParameter bool
	for filterParameter, _ := range filters {
		isAllowedParameter = false

		for _, allowedParameter := range allowedParameters {
			if allowedParameter == filterParameter {
				isAllowedParameter = true
				break
			}
		}

		if !isAllowedParameter {
			return &ValidationError{
				http.StatusUnprocessableEntity,
				fmt.Sprintf(
					"Não existe um filtro para o parâmetro '%s'.",
					filterParameter,
				),
			}
		}
	}

	return nil
}

func ValidateInvalidFilterValues(filters map[string]string) *ValidationError {
	for filterParameter, filterValue := range filters {
		for _, exactNumberParameter := range ExactNumberParameters {
			if exactNumberParameter == filterParameter {
				if _, err := strconv.Atoi(filterValue); err != nil {
					return &ValidationError{
						http.StatusUnprocessableEntity,
						fmt.Sprintf(
							"O valor '%s' é inválido para o parâmetro '%s'.",
							filterValue,
							filterParameter,
						),
					}
				}
			}
		}

		for _, comparativeNumberParameters := range ComparativeNumberParameters {
			if comparativeNumberParameters == filterParameter {
				if len(filterValue) > 0 {
					if filterValue[0:1] == "<" || filterValue[0:1] == ">" {
						if len(filterValue) > 1 && filterValue[1:2] == "=" {
							filterValue = filterValue[2:]
						} else {
							filterValue = filterValue[1:]
						}
					}
				}
				if _, err := strconv.Atoi(filterValue); err != nil {
					return &ValidationError{
						http.StatusUnprocessableEntity,
						fmt.Sprintf(
							"O valor '%s' é inválido para o parâmetro '%s'.",
							filterValue,
							filterParameter,
						),
					}
				}
			}
		}
	}

	return nil
}

func ValidateSuperExists(uuid string) *ValidationError {
	if super := GetSuperByUUID(uuid); super == nil {
		return &ValidationError{
			http.StatusNotFound,
			fmt.Sprintf("O super '%s' não existe.", uuid),
		}
	}

	return nil
}
