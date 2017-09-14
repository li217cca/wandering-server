package model

import "image"

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Town struct {
	ID       int         `json:"id"`
	Name     string      `json:"name";gorm:"not null;unique"`
	Position image.Point `json:"position"`
}
