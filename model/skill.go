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

	Level     int     `json:"level,omitempty"`
	Exp       int     `json:"exp,omitempty"`
	ExpLimit  int     `json:"exp_limit,omitempty"`
	Potential float64 `json:"potential,omitempty"`
}

// SkillWeight ...
type SkillWeight struct {
	Attack  int
	Speed   int
	Defence int
	Hinder  int
	Support int
}

// Skill const
const (
	// extra
	SkillBagWeightBaseID   = 10100 // 基础负重, 1 per level
	SkillBagCapacityBaseID = 10110 // 基础打包, 1 per level

	// special defence
	SkillAttentionDefenceID = 700 // 专注防御, 防御增益 1% per level
	SkillDebuffDecenceID    = 710 // 异常耐受, level/(level+50) * 100% 耐受

	// special attack
	SkillMultiAttackID   = 800 // 追加攻击, 1% per level
	SkillFieldAttackID   = 810 // 全域攻击, level/(level+50) * 100%
	SkillCritAttackID    = 820 // 致命攻击, level/(level+50) * 100%
	SkillRemoteAttackID  = 830 // 远程攻击, 50% * level/(level+50) * 50%
	SkillCounterAttackID = 850 // 反击, level/(level+50) * 100%

	// damage calc
	SkillCritDamageIncreaseID    = 925 // 致命伤害增幅, 1% per level
	SkillCruelDamageIncreaseID   = 930 // 残忍伤害增幅, 100% * (1-hpp) * 1% per level
	SkillCounterDamageIncreaseID = 940 // 反击伤害增幅, 1% per level

	// support
	SkillPurifyTargetID = 2200 // 目标净化, 5% per level
	SkillPurifyFieldID  = 2205 // 领域净化, 2% per level
	SkillBariaTargetID  = 2210 // 目标屏障, 5% per level
	SkillBariaFieldID   = 2205 // 领域屏障, 2% per level

	// special support
	SkillMagicTargetID = 2300 // 目标魔法, 1% per level
	SkillMagicFieldID  = 2305 // 领域魔法, 1% per level

	// hinder
	SkillPoisonTargetID = 3200 // 目标猛毒, 5% per level
	SkillPoisonAttackID = 3210 // 攻击猛毒, 5% per level
	SkillPoisonFieldID  = 3220 // 领域猛毒, 2% per level

)

// GetWeight [pure]
func (skill *Skill) GetWeight() SkillWeight {
	switch skill.Type {
	case SkillAttentionDefenceID:
		return SkillWeight{Defence: 90, Support: 10}
	case SkillBariaFieldID:
		return SkillWeight{Support: 80, Defence: 20}
	case SkillBariaTargetID:
		return SkillWeight{Support: 80, Defence: 20}
	case SkillCounterAttackID:
		return SkillWeight{Attack: 50, Defence: 50}
	case SkillCounterDamageIncreaseID:
		return SkillWeight{Attack: 50, Defence: 50}
	case SkillCritAttackID:
		return SkillWeight{Attack: 80, Speed: 20}
	case SkillCritDamageIncreaseID:
		return SkillWeight{Attack: 80, Speed: 20}
	case SkillCruelDamageIncreaseID:
		return SkillWeight{Attack: 80, Hinder: 20}
	case SkillDebuffDecenceID:
		return SkillWeight{Defence: 80, Support: 10, Attack: 10}
	case SkillFieldAttackID:
		return SkillWeight{Attack: 100}
	case SkillMultiAttackID:
		return SkillWeight{Attack: 90, Speed: 10}
	case SkillPoisonAttackID:
		return SkillWeight{Attack: 30, Hinder: 70}
	case SkillPoisonFieldID:
		return SkillWeight{Hinder: 80, Support: 20}
	case SkillPoisonFieldID:
		return SkillWeight{Hinder: 90, Support: 10}
	case SkillPurifyFieldID:
		return SkillWeight{Defence: 20, Support: 80}
	case SkillPurifyTargetID:
		return SkillWeight{Defence: 10, Support: 90}
	case SkillRemoteAttackID:
		return SkillWeight{Attack: 50, Speed: 50}
	}

	return SkillWeight{}
}

// TypeName ...
func (skill *Skill) TypeName() string {
	switch skill.Type {
	case SkillBagWeightBaseID:
		return "基础负重"
	case SkillBagCapacityBaseID:
		return "基础打包"
	case SkillHitPointIncreaseID:
		return "体力增幅"
	case SkillAttackIncreaseID:
		return "攻击增幅"
	}
	return "未知类型技能"
}

// SkillSet ...
type SkillSet []Skill

func (skills SkillSet) preCalcInt(preFunc func(skill *Skill) int) int {
	tot := 0
	for index := range skills {
		tot += preFunc(&skills[index])
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
func (skill *Skill) LevelUp(diff int) {
	skill.Level += diff
	if skill.Level < 0 {
		skill.Level = 0
	}
	skill.calcExpLimit()
}

// addExp [pure]
func (skill *Skill) AddExp(exp int) {
	expDiff := int(float64(exp) * skill.Potential)
	skill.Potential *= math.Pow(0.99, float64(expDiff)/10)
	skill.Exp += expDiff
	for skill.Exp >= skill.ExpLimit {
		skill.Exp -= skill.ExpLimit
		skill.LevelUp(1)
	}
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

// ToString ...
func (skill *Skill) ToString() string {
	str := fmt.Sprintf(`
		Skill{ID: %d, CharID: %d, Name: %s, Type: %s, Level: %d, Exp: %d/%d, Potential: %.1f%%}`,
		skill.ID, skill.CharactorID, skill.Name, skill.TypeName(), skill.Level,
		skill.Exp, skill.ExpLimit, skill.Potential*100)
	return str
}
