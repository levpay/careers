package superheroapi

import (
	"encoding/json"
	"net/http"
)

type APIResponseBody struct {
	Status  string  `json:"status"`
	Supers  []Super `json:"supers"`
	Message string  `json:"message"`
}

type APIResponse struct {
	httpStatus int
	body       APIResponseBody
}

func createResponseError(validationError *ValidationError) APIResponse {
	emptySupers := []Super{}
	responseBody := APIResponseBody{
		"failed",
		emptySupers,
		validationError.message,
	}
	return APIResponse{
		validationError.httpStatus,
		responseBody,
	}
}

func createResponseSucess(supers []Super) APIResponse {
	responseBody := APIResponseBody{"sucess", supers, ""}
	return APIResponse{http.StatusOK, responseBody}
}

func getParameter(request *http.Request, name string) string {
	return request.PostFormValue("name")
}

func writeResponse(responseWriter http.ResponseWriter, response APIResponse) {
	responseWriter.WriteHeader(response.httpStatus)
	json.NewEncoder(responseWriter).Encode(response.body)
}

func AddSuperHandler(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	name := getParameter(request, "name")

	if err := ValidateParameterRequired("name", name); err != nil {
		writeResponse(responseWriter, createResponseError(err))
		return
	}

	response, searchErr := SearchSuperHeroAPI(name)
	defer response.Body.Close()
	if err := ValidateErrorInSuperHeroAPI(response, searchErr); err != nil {
		writeResponse(responseWriter, createResponseError(err))
		return
	}

	superHeroAPIResponse := GetSuperHeroAPIResponseFromResponse(response)

	if err := ValidateSuperExistsInSuperHeroAPI(superHeroAPIResponse, name); err != nil {
		writeResponse(responseWriter, createResponseError(err))
		return
	}

	supers := ConvertSuperHeroAPIResponseToSuper(superHeroAPIResponse)
	AddSupersDatabase(supers)
	writeResponse(responseWriter, createResponseSucess(supers))
}
