package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/carlsonsantana/superheroapi"
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

func TestAPISuperListInvalidParameter(t *testing.T) {
	response := requestPath("GET", "/super", map[string]string{
		"invalidparameter": "invalid",
	})

	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
	}
}

func TestAPISuperListInvalidValue(t *testing.T) {
	response1 := requestPath("GET", "/super", map[string]string{
		"superheroapi-id": ">3",
	})
	if response1.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response2 := requestPath("GET", "/super", map[string]string{
		"superheroapi-id": "3",
	})
	if response2.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response3 := requestPath("GET", "/super", map[string]string{
		"power": "a",
	})
	if response3.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response4 := requestPath("GET", "/super", map[string]string{
		"power": "a",
	})
	if response4.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response5 := requestPath("GET", "/super", map[string]string{
		"power": ">",
	})
	if response5.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response6 := requestPath("GET", "/super", map[string]string{
		"power": ">=",
	})
	if response6.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response7 := requestPath("GET", "/super", map[string]string{
		"power": "",
	})
	if response7.StatusCode != http.StatusUnprocessableEntity {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response8 := requestPath("GET", "/super", map[string]string{
		"power": "3",
	})
	if response8.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response9 := requestPath("GET", "/super", map[string]string{
		"power": ">3",
	})
	if response9.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response10 := requestPath("GET", "/super", map[string]string{
		"power": ">=3",
	})
	if response10.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response11 := requestPath("GET", "/super", map[string]string{
		"power": "<3",
	})
	if response11.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}

	response12 := requestPath("GET", "/super", map[string]string{
		"power": "<=3",
	})
	if response12.StatusCode != http.StatusOK {
		t.Error("O webservice não esta devolvendo o código correto quando o parâmetro name não é informado")
		return
	}
}

func TestAPISuperListFiltered(t *testing.T) {
	requestPath("POST", "/super", map[string]string{
		"name": "flash",
	})

	response2 := requestPath("GET", "/super", map[string]string{
		"name": "noone",
	})
	body2, _ := ioutil.ReadAll(response2.Body)
	apiResponseBody2 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body2, apiResponseBody2)
	if len(apiResponseBody2.Supers) != 0 {
		t.Error("O não esta filtrando como deveria")
		return
	}

	response3 := requestPath("GET", "/super", map[string]string{
		"name": "black flash",
	})
	body3, _ := ioutil.ReadAll(response3.Body)
	apiResponseBody3 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body3, apiResponseBody3)
	if len(apiResponseBody3.Supers) != 1 {
		t.Error("O não esta filtrando como deveria")
		return
	}

	response4 := requestPath("GET", "/super", map[string]string{
		"name": "%flash%",
	})
	body4, _ := ioutil.ReadAll(response4.Body)
	apiResponseBody4 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body4, apiResponseBody4)
	if len(apiResponseBody4.Supers) != 8 {
		t.Error("O não esta filtrando como deveria")
		return
	}

	response5 := requestPath("GET", "/super", map[string]string{
		"name":  "%flash%",
		"power": ">50",
	})
	body5, _ := ioutil.ReadAll(response5.Body)
	apiResponseBody5 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body5, apiResponseBody5)
	if len(apiResponseBody5.Supers) != 5 {
		t.Error("O não esta filtrando como deveria")
		return
	}

	response6 := requestPath("GET", "/super", map[string]string{
		"name":  "%flash%",
		"power": "100",
	})
	body6, _ := ioutil.ReadAll(response6.Body)
	apiResponseBody6 := &superheroapi.APIResponseBody{}
	json.Unmarshal(body6, apiResponseBody6)
	if len(apiResponseBody6.Supers) != 4 {
		t.Error("O não esta filtrando como deveria")
		return
	}
}
