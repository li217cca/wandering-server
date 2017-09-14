package model

import ()

type Character struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id";gorm:"index"`
	Name   string  `json:"name";gorm:"not null"`
	TownID string  `json:"town_id";gorm:"index"`
	Skills []Skill `json:"skills";gorm:"-"`
}

type Party struct {
	Chars []Character
}

func GetPartyByID(ID int) (p Party) {
	DB.Model(Character{}).Where(Character{ID: ID}).Find(&p.Chars)
	for index, _ := range p.Chars {
		DB.Model(Skill{}).Where(Skill{CharID: p.Chars[index].ID}).Find(&p.Chars[index].Skills)
	}
	return p
}
