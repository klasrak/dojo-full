package swapi

import (
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

func (sw *swapiClient) GetStarship(id int) (models.Starship, error) {
	return models.Starship{}, nil
}

func (sw *swapiClient) GetStarships() (models.Starships, error) {
	return models.Starships{}, nil
}

func (sw *swapiClient) GetPeople(id int) (models.People, error) {
	return models.People{}, nil
}

func (sw *swapiClient) GetPeoples() (models.Peoples, error) {
	return models.Peoples{}, nil
}
