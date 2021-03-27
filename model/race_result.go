package model

type RaceResult struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Distance int32 `json:"distance,omitempty"`

	Course string `json:"course,omitempty"`

	Condition string `json:"condition,omitempty"`

	HorseRaceReults []HorseRaceReult `json:"HorseRaceReults,omitempty"`
}
