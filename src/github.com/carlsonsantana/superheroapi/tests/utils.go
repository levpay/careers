package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func createRequest(
	method string,
	path string,
	parameters map[string]string,
) *http.Request {
	data := url.Values{}
	for key, value := range parameters {
		data.Set(key, value)
	}

	request, _ := http.NewRequest(
		method,
		path,
		strings.NewReader(data.Encode()),
	)

	if method == "POST" {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else if method == "GET" {
		query := request.URL.Query()
		for key, value := range parameters {
			query.Set(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}
	return request
}

func requestPath(
	method string,
	path string,
	parameters map[string]string,
) *http.Response {
	request := createRequest(method, path, parameters)
	responseRecorder := httptest.NewRecorder()

	TestRouter.ServeHTTP(responseRecorder, request)
	return responseRecorder.Result()
}
