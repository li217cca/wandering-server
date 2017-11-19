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
	ID     int `json:"id"`
	GameID int `json:"game_id"`
	InTeam int `json:"in_team"`

	HitPoint      int
	HitPointLimit int
	Skills        []Skill `json:"skills" gorm:"-"`
}

/*
Char const
*/
const (
	CharIsInTeam  = 1
	CharNotInTeam = -1
)

/*
Charactor.commitWithoutChildren
Type: not pure
UnitTest: false
*/
func (char *Charactor) commitWithoutChildren() {
	DB.Save(&char)
}

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
Charactor.commit
Type: not pure
UnitTest: false
*/
func (char *Charactor) commit() {
	for index := range char.Skills {
		char.Skills[index].commit()
	}
	char.commitWithoutChildren()
}

func (char *Charactor) delete() error {
	for index := range char.Skills {
		if err := char.Skills[index].delete(); err != nil {
			return fmt.Errorf("Charactor.delete 01\n %v", err)
		}
	}
	if num := DB.Model(char).Where("id = ?", char.ID).Delete(&char).RowsAffected; num != 1 {
		return fmt.Errorf("Charactor.delete 02\n RowsAffected = %d", num)
	}
	return nil
}

/*
NewCharactor New a Charactor{GameID, Inteam, Skills}
Type: not pure
UnitTest: false
*/
func NewCharactor(gameID int, inTeam int, skills []Skill) Charactor {
	char := Charactor{
		GameID:        gameID,
		InTeam:        inTeam,
		HitPoint:      10,
		HitPointLimit: 10,
		Skills:        skills,
	}
	char.refreshHitPoint(skills)
	char.commitWithoutChildren()
	for index := range skills {
		skills[index].CharactorID = char.ID
	}
	char.commit()
	return char
}
