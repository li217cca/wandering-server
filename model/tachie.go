package model

import "image"

type Tachie struct {
	ID              int    `json:"id"`
	URL             string `json:"url"`
	HeadTransform   string `json:"head_transform"`
	BigTransform    string `json:"big_transform"`
	BattleTransform string `json:"battle_transform"`
	CardTransform   string `json:"card_transform"`
}

type TachieTransform struct {
	Size image.Point
	Position image.Point
}