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
