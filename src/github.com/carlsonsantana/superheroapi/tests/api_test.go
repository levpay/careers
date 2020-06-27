package tests

import (
	"encoding/json"
	"github.com/carlsonsantana/superheroapi"
	"io/ioutil"
	"net/http"
	"reflect"
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

func TestAPISuperAdded(t *testing.T) {
	response := requestPath("POST", "/super", map[string]string{
		"name": "batman",
	})

	if response.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando é informado um super que nao existe")
		return
	}
	body, _ := ioutil.ReadAll(response.Body)
	apiResponseBody := &superheroapi.APIResponseBody{}
	json.Unmarshal(body, apiResponseBody)
	if len(apiResponseBody.Supers) == 0 {
		t.Error("O webservice não retornou nenhum super pesquisando por batman")
	}
}

func TestAPISuperNoDuplicate(t *testing.T) {
	response1 := requestPath("POST", "/super", map[string]string{
		"name": "superman",
	})
	body1, _ := ioutil.ReadAll(response1.Body)
	apiResponseBody1 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body1, apiResponseBody1)

	response2 := requestPath("POST", "/super", map[string]string{
		"name": "superman",
	})
	body2, _ := ioutil.ReadAll(response2.Body)
	apiResponseBody2 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body2, apiResponseBody2)

	if !reflect.DeepEqual(apiResponseBody1, apiResponseBody2) {
		t.Error("O webservice esta duplicando os supers cadastrados")
	}
}

func TestAPISuperListAll(t *testing.T) {
	response1 := requestPath("POST", "/super", map[string]string{
		"name": "wonder woman",
	})
	body1, _ := ioutil.ReadAll(response1.Body)
	apiResponseBody1 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body1, apiResponseBody1)

	response2 := requestPath("GET", "/super", map[string]string{})
	body2, _ := ioutil.ReadAll(response2.Body)
	apiResponseBody2 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body2, apiResponseBody2)

	superFound := false
	for _, super1 := range apiResponseBody1.Supers {
		for _, super2 := range apiResponseBody2.Supers {
			if reflect.DeepEqual(super1, super2) {
				superFound = true
				break
			}
		}
		if superFound {
			break
		}
	}
	if !superFound {
		t.Error("O webservice não esta listando os supers cadastrados")
	}
}
