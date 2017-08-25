package model

import ()

type Skill struct {
	ID     int    `json:"id"`
	CharID int    `json:"char_id";gorm:"index"`
	Name   string `json:"name";gorm:"not null;unique"`
	Level  int    `json:"level"`
	Exp    int    `json:"exp"`
}
