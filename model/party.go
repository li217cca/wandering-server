package model

import (
	"fmt"
	"wandering-server/common"
)

// Party ...
type Party struct {
	ID     int          `json:"id,omitempty"`
	GameID int          `json:"game_id,omitempty"`
	Chars  [6]Charactor `json:"chars,omitempty" gorm:"many2many:party_chars"`
}

// PartyCondition ...
type PartyCondition struct {
	ID      int   `json:"id,omitempty"`
	PartyID int   `json:"party_id,omitempty"`
	Party   Party `json:"party,omitempty"`
}

// Party const
var (
	PartyStrength = map[int]int{
		EnemyDifficultyNormalID: 600,
		EnemyDifficultyEliteID:  1200,
		EnemyDifficultyBossID:   1800,
	}
	PartyRoulette = map[int]map[int]common.Roulette{
		EnemyDifficultyNormalID: {
			EnemyPlantID: common.Roulette{
				{200, CharTypeAttackID},
				{60, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{200, CharTypeHinderID},
			},
			EnemyPhytozoonID: common.Roulette{
				{150, CharTypeAttackID},
				{150, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeHinderID},
			},
			EnemyCarnivoreID: common.Roulette{
				{200, CharTypeAttackID},
				{50, CharTypeDefenceID},
				{50, CharTypeSupportID},
			},
			EnemyCivilizationID: common.Roulette{
				{100, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeHinderID},
			},
			EnemyCivilizationExileID: common.Roulette{
				{60, CharTypeAttackID},
				{60, CharTypeDefenceID},
				{200, CharTypeRemoteAttackID},
				{200, CharTypeHinderID},
			},
			EnemyLegendID: common.Roulette{
				{200, CharTypeDefenceID},
				{50, CharTypeBossID},
			}},
		EnemyDifficultyEliteID: {
			EnemyPlantID: common.Roulette{
				{50, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{30, CharTypeBossID},
				{200, CharTypeHinderID},
			},
			EnemyPhytozoonID: common.Roulette{
				{200, CharTypeAssaultID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{10, CharTypeBossID},
			},
			EnemyCarnivoreID: common.Roulette{
				{200, CharTypeAssaultID},
				{50, CharTypeDefenceID},
				{50, CharTypeHinderID},
				{30, CharTypeSupportID},
				{10, CharTypeBossID},
			},
			EnemyCivilizationID: common.Roulette{
				{50, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{30, CharTypeBossID},
				{30, CharTypeHinderID},
			},
			EnemyCivilizationExileID: common.Roulette{
				{200, CharTypeAssaultID},
				{200, CharTypeRemoteAttackID},
				{200, CharTypeHinderID},
			},
			EnemyLegendID: common.Roulette{
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeBossID},
			},
		},
		EnemyDifficultyBossID: {
			EnemyPlantID: common.Roulette{
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
				{100, CharTypeHinderID},
			},
			EnemyPhytozoonID: common.Roulette{
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
			},
			EnemyCarnivoreID: common.Roulette{
				{50, CharTypeHinderID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
			},
			EnemyCivilizationID: common.Roulette{
				{100, CharTypeRemoteAttackID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
				{50, CharTypeHinderID},
			},
			EnemyCivilizationExileID: common.Roulette{
				{100, CharTypeHinderID},
				{150, CharTypeBossID},
				{50, CharTypeSupportID},
			},
			EnemyLegendID: common.Roulette{
				{70, CharTypeSupportID},
				{70, CharTypeAttackID},
				{150, CharTypeBossID},
			},
		},
	}
)

// NewPartyFromQuest [pure]
func NewPartyFromQuest(danger int, enemyType int, themeType int, difficulty int) Party {
	var chars [6]Charactor

	rou, ok := PartyRoulette[difficulty][enemyType]
	if !ok {
		return Party{}
	}
	length := common.Float(1, 6)
	strength := PartyStrength[difficulty]
	strLimit := strength * 2 / length
	for i := 0; i < length; i++ {
		charType := rou.Get().(int)
		if charType != 0 {
			ranStr := common.Float(0, strLimit)
			if ranStr > strength {
				ranStr = strength
			}
			if ranStr < 50 {
				continue
			}
			strength -= ranStr
			chars[i] = NewCharactorFromQuest(danger, ranStr, enemyType, themeType, charType)
		}
	}
	party := Party{
		Chars: chars,
	}
	return party
}

// ToString [pure]
func (party *Party) ToString() string {
	str := fmt.Sprintf(`
		Party{ID: %d, GameID: %d}`, party.ID, party.GameID)
	for index := range party.Chars {
		str += party.Chars[index].ToString()
	}
	return str
}
