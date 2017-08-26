package model

type Card struct {
	Name string `json:"name"`
	TachieID int `json:"tachie_id"`
	Type string `json:"type"`
	Rare int `json:"rare"`
}

type Scene struct {
	Cards []Card `json:"cards"`
}
