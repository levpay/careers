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
