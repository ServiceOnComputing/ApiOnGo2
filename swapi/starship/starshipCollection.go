package starship

import (
	"ApiOnGo/swapi/collection"
	"ApiOnGo/swapi/errors"
)

type Starship struct {
	errors.ErrorReport

	MGLT                 string   `json:"MGLT"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	CostInCredits        string   `json:"cost_in_credits"`
	Created              string   `json:"created"`
	Crew                 string   `json:"crew"`
	Edited               string   `json:"edited"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	Length               string   `json:"length"`
	Manufacturer         string   `json:"manufacturer"`
	MaxAtmospheringSpeed string   `json:"max_atmostphering_speed"`
	Model                string   `json:"model"`
	Name                 string   `json:"name"`
	Passengers           string   `json:"passengers"`
	Films                []string `json:"films"`
	Pilots               []string `json:"pilots"`
	StarshipClass        string   `json:"starship_class"`
	Url                  string   `json:"url"`
}

type StarshipCollection struct {
	errors.ErrorReport
	collection.ResultCollection
	Results []Starship `json:"results"`
}
