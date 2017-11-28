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
	Name   string  `json:"name,omitempty"    gorm:"not null"`
	Skills []Skill `json:"skills,omitempty"`
}

/*
Char const
*/
const (
	CharTypeAttackID      = 100
	CharTypeSpeedAttackID = 101
	CharTypeDefenceID     = 200
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

// NewCharactorFromQuest [pure]
func NewCharactorFromQuest(danger int, enemyType int, themeType int, difficulty int, charType int) Charactor {
	char := Charactor{}
	// TODO: finish it
	return char
}
