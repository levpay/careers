package tests

import (
	"net/http"
	"testing"
)

func TestAPISuperParameterRequired(t *testing.T) {
	response := requestPath("POST", "/super", map[string]string{})

	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
	}
}

func TestAPISuperNotFoundAPI(t *testing.T) {
	response := requestPath("POST", "/super", map[string]string{
		"name": "supernaoencontrado",
	})

	if response.StatusCode != http.StatusFailedDependency {
		t.Error("O webservice não esta devolvendo o código correto quando é informado um super que nao existe")
	}
}
