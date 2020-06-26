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

func writeResponseErrorIfHasError(
	responseWriter http.ResponseWriter,
	validationError *ValidationError,
) {
	if validationError != nil {
		writeResponse(responseWriter, createResponseError(validationError))
	}
}

func AddSuperHandler(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	name := getParameter(request, "name")

	writeResponseErrorIfHasError(
		responseWriter,
		ValidateParameterRequired("name", name),
	)
	writeResponseErrorIfHasError(
		responseWriter,
		ValidateSuperExistsInAPI(name),
	)
	response, _ := SearchSuperHeroAPI(name)
	superHeroAPIResponse := GetSuperHeroAPIResponseFromResponse(response)
	supers := ConvertSuperHeroAPIResponseToSuper(superHeroAPIResponse)
	AddSupersDatabase(supers)
	writeResponse(responseWriter, createResponseSucess(supers))
}
