package model

import (
)

type Skill struct {
	ID     int    `json:"id"`
	CharID int    `json:"char_id";gorm:"index"`

	Name   string `json:"name";gorm:"not null;unique"`
	Level  int    `json:"level"`
	Exp    int    `json:"exp"`

	ExpLimit int `json:"exp_limit";gorm:"-"`
}

func GetSkillsByCharID(CharID int) (skills []Skill) {
	DB.Model(Skill{}).Where("char_id = ?", CharID).Find(&skills)
	return skills
}

func (s *Skill) Commit() {
	DB.Model(Skill{}).Where("id = ?", s.ID).Update(&s)
}