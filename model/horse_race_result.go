package model

type HorseRaceReult struct {
	Arrival int32 `json:"arrival,omitempty"`
	// 枠番
	BracketNumber int32 `json:"bracket Number,omitempty"`
	// 馬番
	HorseNumber int32 `json:"horse number,omitempty"`

	Horse *Horse `json:"Horse,omitempty"`
	// 斤量
	Weight float64 `json:"weight,omitempty"`
	// 騎手
	Jockey string `json:"jockey,omitempty"`

	Time string `json:"time,omitempty"`

	Difference string `json:"difference,omitempty"`

	Pass string `json:"pass,omitempty"`

	Last3F float64 `json:"last3F,omitempty"`
	// オッズ
	Odds float64 `json:"odds,omitempty"`
	// 人気
	OrderOfFavariteNumber int32 `json:"order of Favarite Number,omitempty"`
	// 馬体重
	HorseWeight string `json:"horse weight,omitempty"`
}
