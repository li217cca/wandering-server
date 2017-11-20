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
	ID     int    `json:"id"`
	GameID int    `json:"game_id"`
	Name   string `gorm:"not null"`
	InTeam int    `json:"in_team" gorm:"not null"`

	HitPoint      int
	HitPointLimit int
	Skills        []Skill
}

/*
Char const
*/
const (
	CharIsInTeam  = 1
	CharNotInTeam = -1
)

/*
Charactor.refreshHitPoint
Type: pure
UnitTest: true
*/
func (char *Charactor) refreshHitPoint(Skills []Skill) {
	tmpLimit := char.HitPointLimit
	tmpValue := char.HitPoint
	char.HitPointLimit = 10
	for index := range Skills {
		char.HitPointLimit += Skills[index].preCalcHitPoint()
	}
	char.HitPoint = tmpValue * char.HitPointLimit / tmpLimit
}

/*
NewCharactor New a Charactor{GameID, Inteam, Skills}
Type: pure
UnitTest: false
*/
func NewCharactor(gameID int, name string, inTeam int) Charactor {
	char := Charactor{
		GameID:        gameID,
		Name:          name,
		InTeam:        inTeam,
		HitPoint:      10,
		HitPointLimit: 10,
	}
	return char
}
