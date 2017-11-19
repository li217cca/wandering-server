package model

import (
	"fmt"
)

// Skill ...
type Skill struct {
	ID          int    `json:"id"`
	CharactorID int    `json:"charactor_id"`
	Name        string `json:"name"    gorm:"not null;unique"`
	Type        int    `json:"type"    gorm:"not null"`

	Level     int `json:"level"`
	Exp       int
	ExpLimit  int
	Potential float64 `json:"potential"`
}

// Skill const
const (
	SkillBagWeightID   = 100
	SkillBagCapacityID = 110
	SkillHitPointID    = 200
)

func (skill *Skill) commit() {
	DB.Where("id = ?", skill.ID).FirstOrCreate(&skill)
	DB.Model(skill).Update(&skill)
}
func (skill *Skill) delete() error {
	if num := DB.Where("id = ?", skill.ID).Delete(&skill).RowsAffected; num != 1 {
		err := fmt.Errorf("RowsAffected = %d", num)
		return fmt.Errorf("Skill.delete 01\n %v", err)
	}
	return nil
}

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
	skill.commit()
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

// NewSkill commited
func NewSkill(Name string, TypeID int, Level int) Skill {
	skill := Skill{
		Name:      Name,
		Type:      TypeID,
		Level:     Level,
		Potential: 1,
	}
	skill.calcExpLimit()
	skill.commit()
	return skill
}

// NewBodySkill commited
func NewBodySkill(Name string, Level int) Skill {
	return NewSkill(Name, SkillHitPointID, Level)
}

// NewCapacitySkill commited
func NewCapacitySkill(Name string, Level int) Skill {
	return NewSkill(Name, SkillBagCapacityID, Level)
}

// NewWeightSkill commited
func NewWeightSkill(Name string, Level int) Skill {
	return NewSkill(Name, SkillBagWeightID, Level)
}
