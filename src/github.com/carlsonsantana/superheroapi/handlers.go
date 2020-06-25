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

func AddSuper(responseWriter http.ResponseWriter, request *http.Request) {
	name := getParameter(request, "name")

	writeResponseErrorIfHasError(
		responseWriter,
		ValidateParameterRequired("name", name),
	)
	writeResponseErrorIfHasError(
		responseWriter,
		ValidateSuperExistsInAPI(name),
	)
}
