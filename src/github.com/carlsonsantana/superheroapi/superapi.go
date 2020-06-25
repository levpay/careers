package superheroapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SearchSuper(name string) (*http.Response, error) {
	url := fmt.Sprintf(
		"https://www.superheroapi.com/api/%s/search/%s",
		os.Getenv("SUPERHEROAPI_TOKEN"),
		name,
	)
	return http.Get(url)
}

func GetSuperAPIResponseFromResponse(
	response *http.Response,
) *SuperAPIResponse {
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	superAPIResponse := &SuperAPIResponse{}
	json.Unmarshal(body, superAPIResponse)
	return superAPIResponse
}
