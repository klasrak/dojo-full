package api

import (
	"fmt"
	"net/http"
	"swapi/clients/swapi"
	"swapi/errors"
	"swapi/mockeable"
	"swapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStarshipHandler(t *testing.T) {

	type TestCase struct {
		Name                        string
		ID                          interface{}
		ExpectedResponseBody        string
		ExpectedStatusCode          int
		ExpectedMockSuccessResponse models.Starship
		ExpectedMockErrorResponse   error
		ExpectedMockCallCount       int
	}

	testCases := []TestCase{
		{
			Name:               "Success",
			ID:                 9,
			ExpectedStatusCode: http.StatusOK,
			ExpectedMockSuccessResponse: models.Starship{
				Name:                 "Death Star",
				Model:                "DS-1 Orbital Battle Station",
				Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
				CostInCredits:        "1000000000000",
				Length:               "120000",
				MaxAtmospheringSpeed: "n/a",
				Crew:                 "342953",
				Passengers:           "843342",
				CargoCapacity:        "1000000000000",
				Consumables:          "3 years",
				HyperdriveRating:     "4.0",
				MGLT:                 "10",
				Class:                "Deep Space Mobile Battlestation",
				Films: []string{
					"https://swapi.dev/api/films/1/",
				},
			},
			ExpectedResponseBody:  `{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}`,
			ExpectedMockCallCount: 1,
		},
		{
			Name:                 "Bad Request",
			ID:                   "invalid_id",
			ExpectedStatusCode:   http.StatusBadRequest,
			ExpectedResponseBody: `{"type":"BAD_REQUEST","message":"Bad request. Reason: invalid id"}`,
		},
		{
			Name:                      "Not Found",
			ID:                        1,
			ExpectedStatusCode:        http.StatusNotFound,
			ExpectedResponseBody:      `{"type":"NOT_FOUND","message":"resource: starship with id: 1 not found"}`,
			ExpectedMockErrorResponse: errors.NewNotFound("starship", "1"),
			ExpectedMockCallCount:     1,
		},
		{
			Name:                      "Internal Server Error",
			ID:                        1,
			ExpectedStatusCode:        http.StatusInternalServerError,
			ExpectedResponseBody:      `{"type":"INTERNAL_SERVER_ERROR","message":"Internal server error."}`,
			ExpectedMockErrorResponse: errors.NewInternal(),
			ExpectedMockCallCount:     1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// Create client mock
			swapiMock := swapi.MockClient{
				GetStarshipFunc: func(id int) (models.Starship, error) {
					assert.Equal(t, tc.ID, id)

					return tc.ExpectedMockSuccessResponse, tc.ExpectedMockErrorResponse
				},
				GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: tc.ExpectedMockCallCount},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			// Create request
			handlerURL := fmt.Sprintf("/api/v1/starship/%v", tc.ID)

			// Do request
			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			// Assert response
			assert.Equal(t, tc.ExpectedStatusCode, response.StatusCode)
			assert.JSONEq(t, tc.ExpectedResponseBody, response.StringBody())
		})
	}
}

func TestGetStarships(t *testing.T) {

	type TestCase struct {
		Name                        string
		ExpectedResponseBody        string
		ExpectedMockSuccessResponse models.Starships
		ExpectedMockErrorResponse   error
		ExpectedMockCallCount       int
		ExpectedStatusCode          int
	}

	testCases := []TestCase{
		{
			Name: "Success",
			ExpectedMockSuccessResponse: models.Starships{
				Count: 1,
				Results: []models.Starship{
					{
						Name:                 "Death Star",
						Model:                "DS-1 Orbital Battle Station",
						Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
						CostInCredits:        "1000000000000",
						Length:               "120000",
						MaxAtmospheringSpeed: "n/a",
						Crew:                 "342953",
						Passengers:           "843342",
						CargoCapacity:        "1000000000000",
						Consumables:          "3 years",
						HyperdriveRating:     "4.0",
						MGLT:                 "10",
						Class:                "Deep Space Mobile Battlestation",
						Films: []string{
							"https://swapi.dev/api/films/1/",
						},
					},
				},
			},
			ExpectedResponseBody:  `{"count":1,"results":[{"name":"Death Star","model":"DS-1 Orbital Battle Station","starship_class":"Deep Space Mobile Battlestation","manufacturer":"Imperial Department of Military Research, Sienar Fleet Systems","cost_in_credits":"1000000000000","length":"120000","crew":"342953","passengers":"843342","max_atmosphering_speed":"n/a","hyperdrive_rating":"4.0","MGLT":"10","cargo_capacity":"1000000000000","consumables":"3 years","films":["https://swapi.dev/api/films/1/"],"pilots":null}]}`,
			ExpectedMockCallCount: 1,
			ExpectedStatusCode:    http.StatusOK,
		},
		{
			Name:                      "Not Found",
			ExpectedMockErrorResponse: errors.NewNotFound("starships", ""),
			ExpectedResponseBody:      `{"type":"NOT_FOUND","message":"resource: starships not found"}`,
			ExpectedMockCallCount:     1,
			ExpectedStatusCode:        http.StatusNotFound,
		},
		{
			Name:                      "Internal Server Error",
			ExpectedMockErrorResponse: errors.NewInternal(),
			ExpectedResponseBody:      `{"type":"INTERNAL_SERVER_ERROR","message":"Internal server error."}`,
			ExpectedMockCallCount:     1,
			ExpectedStatusCode:        http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// Create client mock
			swapiMock := swapi.MockClient{
				GetStarshipsFunc: func() (models.Starships, error) {
					return tc.ExpectedMockSuccessResponse, tc.ExpectedMockErrorResponse
				},
				GetStarshipsFuncControl: mockeable.CallsFuncControl{ExpectedCalls: tc.ExpectedMockCallCount},
			}

			swapiMock.Use()
			defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

			// Create request
			handlerURL := "/api/v1/starships"

			// Do request
			response := DoRequest(http.MethodGet, handlerURL, nil, "")

			// Assert response
			assert.Equal(t, tc.ExpectedStatusCode, response.StatusCode)
			assert.JSONEq(t, tc.ExpectedResponseBody, response.StringBody())
		})
	}
}
