package model

import (
	"fmt"
)

// Skill ...
type Skill struct {
	ID          int    `json:"id,omitempty"`
	CharactorID int    `json:"charactor_id,omitempty"`
	Name        string `json:"name,omitempty" gorm:"not null;unique"`
	Type        int    `json:"type,omitempty" gorm:"not null"`

	Level     int     `json:"level,omitempty"`
	Exp       int     `json:"exp,omitempty"`
	ExpLimit  int     `json:"exp_limit,omitempty"`
	Potential float64 `json:"potential,omitempty"`
}

// Skill const
const (
	SkillBagWeightID   = 100
	SkillBagCapacityID = 110
	SkillHitPointID    = 200
)

func (skill *Skill) preCalcBagWeight() float64 {
	switch skill.Type {
	case SkillBagWeightID:
		return float64(skill.Level)
	}
	return 0
}
func (skill *Skill) preCalcBagCapacity() int {
	switch skill.Type {
	case SkillBagCapacityID:
		return skill.Level
	}
	return 0
}
func (skill *Skill) preCalcHitPoint() int {
	switch skill.Type {
	case SkillHitPointID:
		return skill.Level * 10
	}
	return 0
}

// calcExpLimit not commited
func (skill *Skill) calcExpLimit() {
	switch skill.Type {
	case SkillHitPointID:
		skill.ExpLimit = skill.Level * 100
		return
	}
	skill.ExpLimit = skill.Level * 100
}

// levelUp commited
func (skill *Skill) levelUp(diff int) {
	skill.Level += diff
	if skill.Level < 0 {
		skill.Level = 0
	}
	skill.calcExpLimit()
}

// addExp commited
func (skill *Skill) addExp(exp int) (err error) {
	if exp == 0 {
		return nil
	}
	diff, err := limitAddInt(&skill.Exp, &skill.ExpLimit, exp)
	if err != nil {
		return fmt.Errorf("Skill.addExp 01\n %v", err)
	}
	skill.levelUp(diff)
	return nil
}

/*
NewSkill New a skill
*/
func NewSkill(Name string, TypeID int, Level int) Skill {
	skill := Skill{
		Name:      Name,
		Type:      TypeID,
		Level:     Level,
		Potential: 1,
	}
	skill.calcExpLimit()
	return skill
}
