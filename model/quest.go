package model

import (
	"wandering-server/common"
)

/*
Quest ...
*/
type Quest struct {
	ID     int     `json:"id,omitempty"`
	PreID  int     `json:"pre_id,omitempty"` // 上一个ID
	MapID  int     `json:"map_id,omitempty" gorm:"not null"`
	Danger float64 `json:"danger,omitempty" gorm:"not null"`

	// Name string `json:"name,omitempty"    gorm:"not null"`
	QuestType  int `gorm:"not null"` // 种类...
	EnemyType  int `gorm:"not null"` // 植物/肉/素食动物/文明/传说...
	ThemeType  int `gorm:"not null"` // 普通/科学/魔法/朋克/灾祸...
	Difficulty int `gorm:"not null"` // 普通/精英/BOSS

	Destiny int `json:"destiny,omitempty"` // 命运
}

// Quest const ...
const (
	EnemyDifficultyNormalID = 100
	EnemyDifficultyEliteID  = 200
	EnemyDifficultyBossID   = 300
	EnemyDifficultyPrizeID  = 400 // 奖励
	EnemyDifficultyFriendID = 500 // 友方

	EnemyThemeNormalID   = 100 // 普通
	EnemyThemeMagicID    = 200 // 魔法
	EnemyThemeScienceID  = 300 // 科学
	EnemyThemePunkID     = 400 // 朋克
	EnemyThemeDisasterID = 500 // 灾祸

	EnemyPlantID             = 100 // 植物
	EnemyPhytozoonID         = 200 // 素食动物
	EnemyCarnivoreID         = 300 // 肉食动物
	EnemyCivilizationID      = 400 // 文明
	EnemyCivilizationExileID = 410 // 放逐文明
	EnemyLegengID            = 500 // 传说

	QuestBattleID         = 200 // [1,   2] 一种种族，概率Prize
	QuestRaidID           = 300 // [3,   9] 一种种族，概率Prize，概率BOSS
	QuestMixedRaidID      = 310 // [4,  12] 几种种族，概率Prize，概率BOSS
	QuestDungeonID        = 400 // [9,  36] 一种种族，少量Prize，少量BOSS
	QuestMixedDungeonID   = 410 // [10, 40] 几种种族，少量Prize，少量BOSS
	QuestChaosDungeonID   = 430 // [5,  25] 几种Elite/BOSS，Prize
	QuestAncientDungeonID = 450 // [X,   X] 远古地牢 // TODO
)

// DestinyResult ...
func (quest *Quest) DestinyResult() {
	// TODO:
}

// GenerateNextQuest [pure]
func (pre *Quest) GenerateNextQuest(lucky float64) Quest {
	quest := Quest{
		PreID:      pre.ID,
		Danger:     common.FloatF(pre.Danger-1, pre.Danger+2),
		QuestType:  0,                                 // TODO
		EnemyType:  0,                                 // TODO
		ThemeType:  0,                                 // TODO
		Difficulty: 0,                                 // TODO
		Destiny:    pre.Destiny + pre.DestinyResult(), // TODO
	}
	return quest
}

// GenerateNextQuest [pure]
func (mp *Map) GenerateNextQuest(lucky float64, pre Quest) {
	rou := common.Roulette{
		{
			Weight: 100,
			Target: QuestNormalBattleID,
		},
		{
			Weight: 3 * pre.Deep,
			Target: QuestNormalBattleID,
		},
		{
			Weight: common.Sqrt(lucky),
			Target: QuestNormalTreasureID,
		},
	}
	quest := Quest{
		PreID: pre.ID,
		Name:  common.GenerateKey(8), // TODO: 生成名称
		Type:  rou.Get().(int),
		Level: common.Float(pre.Level-2, pre.Level+4),
		Deep:  pre.Deep + 1,
	}
	mp.Quests = append(mp.Quests, quest)
}

/*
Map.generateQuest [pure] Generate quest from lucky, danger, miracle ...
*/
func (mp *Map) GenerateQuest(lucky int) {
	mp.GenerateNextQuest(lucky, Quest{
		ID:    0,
		Level: mp.Resource.Danger,
		Deep:  0,
	})
}
