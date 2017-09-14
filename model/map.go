package model

import (
	"wandering-server/model/math"
)

type Map struct {
	ID       int           `json:"id"`
	SpiritID int           `json:"-"`

	Borders  []math.Vector `json:"borders";gorm:"-"`
}

type Line struct {
	ID       int `json:"id"`
	ParentID int `json:"-";gorm:"index;not null"`
	math.Vector
}
