package model

type Horse struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Trainer string `json:"trainer,omitempty"`

	Owner string `json:"owner,omitempty"`

	Breeder string `json:"breeder,omitempty"`

	Sire string `json:"sire,omitempty"`
	// 母父
	Broodmare string `json:"broodmare,omitempty"`
}
