package models

type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Class                string   `json:"starship_class"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	Films                []string `json:"films"`
	Pilots               []string `json:"pilots"`
}

type Starships struct {
	Count   int        `json:"count"`
	Results []Starship `json:"results"`
}

type People struct {
	Name      string   `json:"name"`
	BirthYear string   `json:"birth_year"`
	EyeColor  string   `json:"eye_color"`
	Gender    string   `json:"gender"`
	HairColor string   `json:"hair_color"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	SkinColor string   `json:"skin_color"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Starships []string `json:"starships"`
}

type Peoples struct {
	Count   int      `json:"count"`
	Results []People `json:"results"`
}
