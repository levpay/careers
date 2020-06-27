package superheroapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type SuperHeroAPIResponse struct {
	Response string              `json:"response"`
	Results  []SuperHeroAPISuper `json:"results"`
	Error    string              `json:"error"`
}

func SearchSuperHeroAPI(name string) (*http.Response, error) {
	url := fmt.Sprintf(
		"https://www.superheroapi.com/api/%s/search/%s",
		os.Getenv("SUPERHEROAPI_TOKEN"),
		name,
	)
	return http.Get(url)
}

func GetSuperHeroAPIResponseFromResponse(
	response *http.Response,
) *SuperHeroAPIResponse {
	body, _ := ioutil.ReadAll(response.Body)
	superHeroAPIResponse := &SuperHeroAPIResponse{}
	json.Unmarshal(body, superHeroAPIResponse)
	return superHeroAPIResponse
}
