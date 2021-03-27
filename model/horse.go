package model

type Horse struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Age int32 `json:"age,omitempty"`

	Sex string `json:"sex,omitempty"`

	Trainer string `json:"trainer,omitempty"`

	Owner string `json:"owner,omitempty"`
}
