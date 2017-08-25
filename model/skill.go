package model

import ()

type Skill struct {
	ID     int
	CharID int		`gorm:"index"`
	Name   string `gorm:"not null;unique"`
	Level  int
	Exp    int
}
