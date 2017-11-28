package model

import (
	"fmt"
	"math"
)

// Skill ...
type Skill struct {
	ID          int    `json:"id,omitempty"`
	CharactorID int    `json:"charactor_id,omitempty"`
	Name        string `json:"name,omitempty" gorm:"not null;"`
	Type        int    `json:"type,omitempty" gorm:"not null"`
	ForTeam     int    `json:"for_team,omitempty"`

	Level     int     `json:"level,omitempty"`
	Exp       int     `json:"exp,omitempty"`
	ExpLimit  int     `json:"exp_limit,omitempty"`
	Potential float64 `json:"potential,omitempty"`
}

// Skill const
const (
	SkillForTeam           = 1
	SkillNotForTeam        = -1
	SkillBagWeightBaseID   = 100 // 基础负重, 1 per level
	SkillBagCapacityBaseID = 110 // 基础打包, 1 per level

	SkillHitPointBaseID     = 200 // 基础体力, 1 per Level
	SkillHitPointIncreaseID = 201 // 体力增幅, 1% per level
	SkillAttackBaseID       = 300 // 基础攻击, 1 per level
	SkillAttackIncreaseID   = 301 // 攻击增幅, 1% per
	SkillMultiAttackID      = 310 // 多次攻击
	SkillDefenceBaseID      = 400 // 基础防御, 1 per level
)

// SkillSet ...
type SkillSet []Skill

func (skills SkillSet) preCalcInt(preFunc func(skill *Skill) int) int {
	tot := 0
	for index := range skills {
		tot += preFunc(&skills[index])
	}
	return tot
}
func (skills SkillSet) preCalcTeamInt(preFunc func(skill *Skill) int) int {
	tot := 0
	for index := range skills {
		if skills[index].ForTeam == SkillForTeam {
			tot += preFunc(&skills[index])
		}
	}
	return tot
}
func (skills SkillSet) preCalcTeamFloat64(preFunc func(skill *Skill) float64) float64 {
	tot := 0.
	for index := range skills {
		if skills[index].ForTeam == SkillForTeam {
			tot += preFunc(&skills[index])
		}
	}
	return tot
}

func (skill *Skill) preCalcBagWeightBase() float64 {
	switch skill.Type {
	case SkillBagWeightBaseID:
		return float64(skill.Level)
	}
	return 0
}
func (skill *Skill) preCalcBagCapacityBase() int {
	switch skill.Type {
	case SkillBagCapacityBaseID:
		return skill.Level
	}
	return 0
}
func (skill *Skill) preCalcAttackBase() int {
	switch skill.Type {
	case SkillAttackBaseID:
		return skill.Level
	}
	return 0
}
func (skill *Skill) preCalcAttackIncrease() float64 {
	switch skill.Type {
	case SkillAttackIncreaseID:
		return float64(skill.Level) * 0.01
	}
	return 0
}
func (skill *Skill) preCalcHitpoint() int {
	switch skill.Type {
	case SkillHitPointBaseID:
		return skill.Level
	}
	return 0
}
func (skill *Skill) preCalcHitPointIncrease() float64 {
	switch skill.Type {
	case SkillHitPointIncreaseID:
		return float64(skill.Level) * 0.01
	}
	return 0
}

// calcExpLimit not commited
func (skill *Skill) calcExpLimit() {
	switch skill.Type {

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

// addExp [pure]
func (skill *Skill) addExp(exp int) (err error) {
	if exp == 0 {
		return nil
	}
	expDiff := int(float64(exp) * skill.Potential)
	diff, err := limitAddInt(&skill.Exp, &skill.ExpLimit, expDiff)
	if err != nil {
		return fmt.Errorf("Skill.addExp 01\n %v", err)
	}
	skill.Potential *= math.Pow(0.99, float64(expDiff)/10)
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
