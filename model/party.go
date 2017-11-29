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
		QuestDifficultyDefaultID: 0,
		QuestDifficultyNormalID:  600,
		QuestDifficultyEliteID:   1200,
		QuestDifficultyBossID:    1800,
	}
	PartyRoulette = map[int]map[int]common.Roulette{
		QuestDifficultyDefaultID: {
			EnemyTypeDefaultID: common.Roulette{
				{100, CharTypeDefaultID},
			},
		},
		QuestDifficultyNormalID: {
			EnemyTypeDefaultID: common.Roulette{
				{100, CharTypeDefaultID},
			},
			EnemyTypePlantID: common.Roulette{
				{200, CharTypeAttackID},
				{60, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{200, CharTypeHinderID},
			},
			EnemyTypePhytozoonID: common.Roulette{
				{150, CharTypeAttackID},
				{150, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeHinderID},
			},
			EnemyTypeCarnivoreID: common.Roulette{
				{200, CharTypeAttackID},
				{50, CharTypeDefenceID},
				{50, CharTypeSupportID},
			},
			EnemyTypeCivilizationID: common.Roulette{
				{100, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeHinderID},
			},
			EnemyTypeCivilizationExileID: common.Roulette{
				{60, CharTypeAttackID},
				{60, CharTypeDefenceID},
				{200, CharTypeRemoteAttackID},
				{200, CharTypeHinderID},
			},
			EnemyTypeLegendID: common.Roulette{
				{200, CharTypeDefenceID},
				{50, CharTypeBossID},
			}},
		QuestDifficultyEliteID: {
			EnemyTypeDefaultID: common.Roulette{
				{100, CharTypeDefaultID},
			},
			EnemyTypePlantID: common.Roulette{
				{50, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{30, CharTypeBossID},
				{200, CharTypeHinderID},
			},
			EnemyTypePhytozoonID: common.Roulette{
				{200, CharTypeAssaultID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{10, CharTypeBossID},
			},
			EnemyTypeCarnivoreID: common.Roulette{
				{200, CharTypeAssaultID},
				{50, CharTypeDefenceID},
				{50, CharTypeHinderID},
				{30, CharTypeSupportID},
				{10, CharTypeBossID},
			},
			EnemyTypeCivilizationID: common.Roulette{
				{50, CharTypeAttackID},
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{30, CharTypeBossID},
				{30, CharTypeHinderID},
			},
			EnemyTypeCivilizationExileID: common.Roulette{
				{200, CharTypeAssaultID},
				{200, CharTypeRemoteAttackID},
				{200, CharTypeHinderID},
			},
			EnemyTypeLegendID: common.Roulette{
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{50, CharTypeBossID},
			},
		},
		QuestDifficultyBossID: {
			EnemyTypeDefaultID: common.Roulette{
				{100, CharTypeDefaultID},
			},
			EnemyTypePlantID: common.Roulette{
				{100, CharTypeRemoteAttackID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
				{100, CharTypeHinderID},
			},
			EnemyTypePhytozoonID: common.Roulette{
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
			},
			EnemyTypeCarnivoreID: common.Roulette{
				{50, CharTypeHinderID},
				{100, CharTypeDefenceID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
			},
			EnemyTypeCivilizationID: common.Roulette{
				{100, CharTypeRemoteAttackID},
				{100, CharTypeSupportID},
				{100, CharTypeBossID},
				{50, CharTypeHinderID},
			},
			EnemyTypeCivilizationExileID: common.Roulette{
				{100, CharTypeHinderID},
				{150, CharTypeBossID},
				{50, CharTypeSupportID},
			},
			EnemyTypeLegendID: common.Roulette{
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

	if _, ok := PartyRoulette[difficulty]; !ok {
		difficulty = QuestDifficultyDefaultID
	}
	if _, ok := PartyRoulette[difficulty][enemyType]; !ok {
		enemyType = EnemyTypeDefaultID
	}
	rou := PartyRoulette[difficulty][enemyType]

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
