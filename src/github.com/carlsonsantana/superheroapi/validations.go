package superheroapi

import (
	"fmt"
	"net/http"
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

func ValidateSuperExistsInAPI(name string) *ValidationError {
	response, err := SearchSuperHeroAPI(name)
	if err != nil || response.StatusCode != http.StatusOK {
		return &ValidationError{http.StatusInternalServerError, "Erro interno"}
	}
	superHeroAPIResponse := GetSuperHeroAPIResponseFromResponse(response)
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
