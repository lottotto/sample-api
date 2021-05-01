package model

type Race struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Field string `json:"field,omitempty"`

	Rotation string `json:"rotation,omitempty"`

	Distance int32 `json:"distance,omitempty"`

	Course string `json:"course,omitempty"`

	Condition string `json:"condition,omitempty"`
}
