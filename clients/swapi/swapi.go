package swapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"swapi/models"
)

func NewSWAPIClient() *swapiClient {
	return &swapiClient{
		client:  &http.Client{},
		baseURL: "https://swapi.dev/api",
	}
}

type swapiClient struct {
	client  *http.Client
	baseURL string
}

func (sw *swapiClient) GetStarship(id int) (result models.Starship, err error) {
	resource := fmt.Sprintf("/starships/%d/", id)
	res, err := sw.client.Get(sw.baseURL + resource)

	if err != nil {
		return result, err
	}

	err = getBody(res, &result)

	if err != nil {
		return result, err
	}

	return result, err
}

func (sw *swapiClient) GetStarships() (result models.Starships, err error) {
	resource := "/starships/"
	res, err := sw.client.Get(sw.baseURL + resource)

	if err != nil {
		return result, err
	}

	err = getBody(res, &result)

	if err != nil {
		return result, err
	}

	return result, err
}

func (sw *swapiClient) GetPeople(id int) (result models.People, err error) {
	resource := fmt.Sprintf("/people/%d/", id)
	res, err := sw.client.Get(sw.baseURL + resource)

	if err != nil {
		return result, err
	}

	err = getBody(res, &result)

	if err != nil {
		return result, err
	}

	return result, err
}

func (sw *swapiClient) GetPeoples() (result models.Peoples, err error) {
	resource := "/people/"
	res, err := sw.client.Get(sw.baseURL + resource)

	if err != nil {
		return result, err
	}

	err = getBody(res, &result)

	if err != nil {
		return result, err
	}

	return result, err
}

func getBody(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)

	if err != nil {
		return err
	}

	return nil
}
