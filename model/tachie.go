package model

type Tachie struct {
	ID              int             `json:"id"`
	URL             string          `json:"url"`
	HeadTransform   TachieTransform `json:"head_transform"`
	BigTransform    TachieTransform `json:"big_transform"`
	BattleTransform TachieTransform `json:"battle_transform"`
	CardTransform   TachieTransform `json:"card_transform"`
}

type TachieTransform struct {
	Size     string `json:"size"`
	Position string `json:"position"`
}
