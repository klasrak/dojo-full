package services

import (
	"net/http"
	"swapi/clients/swapi"
	"swapi/errors"
	"swapi/mockeable"
	"swapi/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStarshipService(t *testing.T) {

	type TestCase struct {
		Name                    string
		IsErrorFlow             bool
		IDToCall                int
		ExpectedSuccessResponse models.Starship
		ExpectedErrorResponse   error
		ExpectedCallCount       int
		ExpectedStatusCode      int
	}

	testCases := []TestCase{
		{
			Name: "Success",
			ExpectedSuccessResponse: models.Starship{
				Name:                 "Naboo fighter",
				Model:                "N-1 starfighter",
				Manufacturer:         "Corellian Engineering Corporation",
				CostInCredits:        "105000",
				Length:               "34",
				MaxAtmospheringSpeed: "1200",
				Crew:                 "1",
				Passengers:           "0",
				CargoCapacity:        "25",
				Consumables:          "1 week",
				HyperdriveRating:     "2.0",
				MGLT:                 "60",
				Pilots: []string{
					"https://swapi.co/api/people/1/",
					"https://swapi.co/api/people/9/",
					"https://swapi.co/api/people/18/",
					"https://swapi.co/api/people/81/",
				},
			},
			ExpectedErrorResponse: nil,
			ExpectedCallCount:     1,
			ExpectedStatusCode:    http.StatusOK,
		},
		{
			Name:                  "Not Found",
			IsErrorFlow:           true,
			ExpectedErrorResponse: errors.NewNotFound("starship", "1"),
			IDToCall:              1,
			ExpectedCallCount:     1,
			ExpectedStatusCode:    http.StatusNotFound,
		},
		{
			Name:                  "Internal Server Error",
			IsErrorFlow:           true,
			ExpectedErrorResponse: errors.NewInternal(),
			IDToCall:              1,
			ExpectedCallCount:     1,
			ExpectedStatusCode:    http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {

		// Create mock client
		swapiMock := swapi.MockClient{
			GetStarshipFunc: func(id int) (models.Starship, error) {
				assert.Equal(t, tc.IDToCall, id)
				return tc.ExpectedSuccessResponse, tc.ExpectedErrorResponse
			},
			GetStarshipFuncControl: mockeable.CallsFuncControl{ExpectedCalls: tc.ExpectedCallCount},
		}

		swapiMock.Use()
		defer mockeable.CleanUpAndAssertControls(t, &swapiMock)

		result, err := GetStarshipService(tc.IDToCall)

		if tc.IsErrorFlow {
			assert.NotNil(t, err)
			assert.Equal(t, tc.ExpectedErrorResponse, err)
			assert.Equal(t, tc.ExpectedStatusCode, errors.Status(err))
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tc.ExpectedSuccessResponse, result)
		}
	}
}
