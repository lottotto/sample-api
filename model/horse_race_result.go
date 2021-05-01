package model

type HorseRaceReults struct {
	RaceId string `json:"raceId,omitempty"`

	HorseID string `json:"horseID,omitempty"`

	Arrival int32 `json:"arrival,omitempty"`
	// 枠番
	BracketNumber int32 `json:"bracket Number,omitempty"`
	// 馬番
	HorseNumber int32 `json:"horse number,omitempty"`
	// 性齢
	SexAndAge string `json:"sex_and_age,omitempty"`
	// 斤量
	Weight float64 `json:"weight,omitempty"`
	// 騎手
	Jockey string `json:"jockey,omitempty"`

	Time string `json:"time,omitempty"`

	Difference string `json:"difference,omitempty"`

	Position string `json:"position,omitempty"`

	Last3F float64 `json:"last3F,omitempty"`
	// オッズ
	Odds float64 `json:"odds,omitempty"`
	// 人気
	Favarite int32 `json:"favarite,omitempty"`
	// 馬体重
	HorseWeight string `json:"horse weight,omitempty"`
	// 調教師
	Trainer string `json:"trainer,omitempty"`
	// 馬主
	Owner string `json:"owner,omitempty"`
}
