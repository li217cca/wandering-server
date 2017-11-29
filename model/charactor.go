package model

import (
	"fmt"
)

/*
limitAddInt
Type: pure
UnitTest: true
*/
func limitAddInt(value *int, limit *int, inc int) (diff int, err error) {
	if *limit == 0 {
		return 0, fmt.Errorf("0 devided by %d of limit = 0", value)
	}
	*value += inc
	diff = 0
	for *value > *limit {
		diff++
		*value -= *limit
	}
	for *value < 0 {
		diff--
		*value += *limit
	}
	return diff, err
}

/*
limitAddFloat64
Type: pure
UnitTest: true
*/
func limitAddFloat64(value *float64, limit *float64, inc float64) (diff int, err error) {
	if *limit == 0 {
		return 0, fmt.Errorf("0 devided by %d of limit = 0", value)
	}
	*value += inc
	diff = 0
	for *value > *limit {
		diff++
		*value -= *limit
	}
	for *value < 0 {
		diff--
		*value += *limit
	}
	return diff, err
}

/*
Charactor Charactor struct
*/
type Charactor struct {
	ID     int     `json:"id,omitempty"`
	Name   string  `json:"name,omitempty" gorm:"not null"`
	Life   int     `json:"life,omitempty"`
	Skills []Skill `json:"skills,omitempty"`
	IsDie  bool    `json:"is_die,omitempty"`
}

/*
Char const
*/
const (
	CharTypeAttackID       = 100 // 近战攻击型
	CharTypeRemoteAttackID = 110 // 远程攻击型
	CharTypeAssaultID      = 120 // 突击型
	CharTypeDefenceID      = 200 // 防御型
	CharTypeSupportID      = 300 // 辅助型
	CharTypeBossID         = 400 // BOSS型，均衡
	CharTypeHinderID       = 500 // 妨碍型
)

/*
Charactor.refreshHitPoint
Type: pure
UnitTest: true
*/
// func (char *Charactor) refreshHitPoint(Skills []Skill) {
// 	tmpLimit := char.HitPointLimit
// 	tmpValue := char.HitPoint
// 	char.HitPointLimit = 10
// 	for index := range Skills {
// 		char.HitPointLimit += Skills[index].preCalcHitPoint()
// 	}
// 	char.HitPoint = tmpValue * char.HitPointLimit / tmpLimit
// }

/*
NewCharactor New a Charactor{GameID, Inteam, Skills}
Type: pure
UnitTest: false
*/
func NewCharactor(name string, skills []Skill) Charactor {
	char := Charactor{
		Name: name,
		// HitPoint:      10,
		// HitPointLimit: 10,
		Skills: skills,
	}
	// char.refreshHitPoint(skills)
	return char
}

var (
	CharSkillWeight = map[int]SkillWeight{
		CharTypeAttackID: SkillWeight{
			Attack:  60,
			Speed:   10,
			Defence: 25,
			Hinder:  0,
			Support: 5,
		},
		CharTypeRemoteAttackID: {
			Attack:  50,
			Speed:   35,
			Defence: 5,
			Hinder:  10,
			Support: 0,
		},
		CharTypeAssaultID: {
			Attack:  70,
			Speed:   30,
			Defence: 0,
			Hinder:  0,
			Support: 0,
		},
		CharTypeDefenceID: {
			Attack:  30,
			Speed:   0,
			Defence: 60,
			Hinder:  0,
			Support: 10,
		},
		CharTypeSupportID: {
			Attack:  10,
			Speed:   0,
			Defence: 30,
			Hinder:  0,
			Support: 60,
		},
		CharTypeHinderID: {
			Attack:  10,
			Speed:   25,
			Defence: 5,
			Hinder:  60,
			Support: 0,
		},
		CharTypeBossID: {
			Attack:  50,
			Speed:   10,
			Defence: 30,
			Hinder:  0,
			Support: 10,
		},
	} // CharType
	EnemyThemeSkillList = map[int]map[int][]int{ // EnemyType / ThemeType
		EnemyPlantID: {
			EnemyThemeNormalID: []int{
				SkillAttentionDefenceID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillPoisonTargetID, SkillPoisonFieldID},
			EnemyThemeMagicID: []int{
				SkillAttentionDefenceID, SkillDebuffDecenceID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillPoisonTargetID, SkillPoisonFieldID},
			EnemyThemeScienceID: []int{
				SkillAttentionDefenceID, SkillDebuffDecenceID,
				SkillRemoteAttackID, SkillMultiAttackID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillPoisonAttackID},
			EnemyThemePunkID: []int{
				SkillAttentionDefenceID, SkillDebuffDecenceID,
				SkillRemoteAttackID, SkillMultiAttackID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillPoisonTargetID, SkillPoisonAttackID, SkillPoisonFieldID},
			EnemyThemeDisasterID: []int{
				SkillCritAttackID, SkillCritDamageIncreaseID,
				SkillRemoteAttackID, SkillMultiAttackID, SkillFieldAttackID,
				SkillPoisonAttackID, SkillPoisonFieldID},
		},
		EnemyPhytozoonID: {
			EnemyThemeNormalID: []int{
				SkillAttentionDefenceID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillCritAttackID, SkillCritDamageIncreaseID,
				SkillPoisonAttackID},
			EnemyThemeMagicID: []int{
				SkillAttentionDefenceID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillCritAttackID, SkillCritDamageIncreaseID,
				SkillPoisonAttackID},
			EnemyThemeScienceID: []int{
				SkillAttentionDefenceID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillCritAttackID, SkillCritDamageIncreaseID,
				SkillPoisonAttackID},
			EnemyThemePunkID: []int{
				SkillAttentionDefenceID,
				SkillRemoteAttackID, SkillMultiAttackID,
				SkillCounterAttackID, SkillCounterDamageIncreaseID,
				SkillPoisonTargetID, SkillPoisonAttackID, SkillPoisonFieldID},
			EnemyThemeDisasterID: []int{
				SkillCritAttackID, SkillCritDamageIncreaseID,
				SkillRemoteAttackID, SkillMultiAttackID, SkillFieldAttackID,
				SkillPoisonAttackID, SkillPoisonFieldID},
		},
		EnemyCarnivoreID:         {},
		EnemyCivilizationID:      {},
		EnemyCivilizationExileID: {},
		EnemyLegendID:            {},
	}
)

// NewCharactorFromQuest [pure]
func NewCharactorFromQuest(danger int, strength int, enemyType int, themeType int, charType int) Charactor {
	var skills []Skill
	char := Charactor{}
	// TODO: finish it
	return char
}

// ToString ...
func (char *Charactor) ToString() string {
	str := fmt.Sprintf(`
		Charactor{ID: %d, Name: %s}`, char.ID, char.Name)
	for index := range char.Skills {
		str += char.Skills[index].ToString()
	}
	return str
}
