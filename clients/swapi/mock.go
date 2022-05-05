package swapi

import (
	"swapi/mockeable"
	"swapi/models"
)

type MockClient struct {
	GetStarshipFunc  func(id int) (models.Starship, error)
	GetStarshipsFunc func() (models.Starships, error)
	GetPeopleFunc    func(id int) (models.People, error)
	GetPeoplesFunc   func() (models.Peoples, error)

	GetStarshipFuncControl  mockeable.CallsFuncControl
	GetStarshipsFuncControl mockeable.CallsFuncControl
	GetPeopleFuncControl    mockeable.CallsFuncControl
	GetPeoplesFuncControl   mockeable.CallsFuncControl
}

func (c *MockClient) GetStarship(id int) (models.Starship, error) {
	c.GetStarshipFuncControl.IncreaseCallCount()

	return c.GetStarshipFunc(id)
}

func (c *MockClient) GetStarships() (models.Starships, error) {
	c.GetStarshipsFuncControl.IncreaseCallCount()

	return c.GetStarshipsFunc()
}

func (c *MockClient) GetPeople(id int) (models.People, error) {
	c.GetPeopleFuncControl.IncreaseCallCount()

	return c.GetPeopleFunc(id)
}

func (c *MockClient) GetPeoples() (models.Peoples, error) {
	c.GetPeoplesFuncControl.IncreaseCallCount()

	return c.GetPeoplesFunc()
}

func (c *MockClient) Use() {
	c.GetStarshipFuncControl.SetFuncName("GetStarship")
	c.GetStarshipsFuncControl.SetFuncName("GetStarships")
	c.GetPeopleFuncControl.SetFuncName("GetPeople")
	c.GetPeoplesFuncControl.SetFuncName("GetPeoples")

	Instance = c
}

func (c *MockClient) CleanUp() {
	Instance = defaultInstance
}

func (c *MockClient) GetFuncControls() []*mockeable.CallsFuncControl {
	return []*mockeable.CallsFuncControl{
		&c.GetStarshipFuncControl,
		&c.GetStarshipsFuncControl,
		&c.GetPeopleFuncControl,
		&c.GetPeoplesFuncControl,
	}
}
