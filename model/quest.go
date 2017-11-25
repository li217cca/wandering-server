package model

import (
	"math"
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

func (quest *Quest) Size() int {
	switch quest.QuestType {
	case QuestBattleID:
		return 2
	case QuestRaidID:
		return 6
	case QuestMixedRaidID:
		return 8
	case QuestDungeonID:
		return 23
	case QuestMixedDungeonID:
		return 25
	case QuestChaosDungeonID:
		return 15
	}
	return 0
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

func randomQuestType() int {
	rou := common.Roulette{
		{1, QuestBattleID},
		{1, QuestRaidID},
		{1, QuestMixedRaidID},
		{1, QuestDungeonID},
		{1, QuestMixedDungeonID},
		{1, QuestChaosDungeonID},
		{0, QuestAncientDungeonID},
	}
	return rou.Get().(int)
}

func (quest *Quest) isMixed() bool {
	switch quest.QuestType {
	case QuestMixedRaidID:
		return true
	case QuestMixedDungeonID:
		return true
	case QuestChaosDungeonID:
		return true
	}
	return false
}

func (quest *Quest) randomNextEnemyType(res *Resource) int {
	notMix := 100.
	if !quest.isMixed() {
		notMix = 1000
	}
	notMix *= math.Sqrt(res.Area() / 2000)
	if quest.EnemyType == 0 {
		notMix = 0
	}
	rou := common.Roulette{
		{int(notMix), quest.EnemyType},
		{int(100 * math.Sqrt(res.PlantResource/10000)), EnemyPlantID},
		{int(100 * math.Sqrt(res.PhytozoonResource/1000)), EnemyPhytozoonID},
		{int(100 * math.Sqrt(res.CarnivoreResource/200)), EnemyCarnivoreID},
		{int(100 * math.Sqrt(res.CivilizationResource/200)), EnemyCivilizationID},
		{int(100 * math.Sqrt(res.CivilizationResource/1000)), EnemyCivilizationExileID},
		{int(100 * math.Sqrt(res.LegendResource/1000)), EnemyLegengID},
	}
	return rou.Get().(int)
}

// randomNextThemeType [pure]
func (quest *Quest) randomNextThemeType(res *Resource) int {
	notMix := 100.
	if !quest.isMixed() {
		notMix = 1000
	}
	notMix *= math.Sqrt(res.Area() / 1000)
	if quest.ThemeType == 0 {
		notMix = 0
	}
	rou := common.Roulette{
		{int(notMix), quest.ThemeType},
		{int(100 * math.Sqrt(res.Area()/1000)), EnemyThemeNormalID},
		{int(100 * math.Sqrt(res.MagicResource/1000)), EnemyThemeMagicID},
		{int(100 * math.Sqrt(res.ScienceResource/1000)), EnemyThemeScienceID},
		{int(100 * math.Sqrt(res.MagicResource/500) * math.Sqrt(res.ScienceResource/500)), EnemyThemePunkID},
		{int(100 * math.Sqrt(res.DisasterResource/1000)), EnemyThemeDisasterID},
	}
	return rou.Get().(int)
}

// randomNextDifficulty [pure]
func (quest *Quest) randomNextDifficulty(lucky float64) int {
	destinyAddition := math.Sqrt(float64(quest.Destiny / 1000))
	luckyAddition := math.Sqrt(lucky / 50)
	sizeAddition := math.Sqrt(float64(quest.Size()) / 10)
	rou := common.Roulette{
		{int(100), EnemyDifficultyNormalID},
		{int(5 + 50*destinyAddition), EnemyDifficultyEliteID},
		{int(50 * destinyAddition * sizeAddition), EnemyDifficultyBossID},
		{int(5 + 50*destinyAddition*luckyAddition), EnemyDifficultyPrizeID},
		{int(50 * luckyAddition * sizeAddition), EnemyDifficultyFriendID},
	}
	return rou.Get().(int)
}

// GenerateNextQuest [pure]
func (mp *Map) GenerateNextQuest(lucky float64, destinyDiff int, pre *Quest) Quest {
	quest := Quest{
		PreID:      pre.ID,
		Danger:     common.FloatF(pre.Danger-1, pre.Danger+2),
		QuestType:  pre.QuestType,                         // TODO
		EnemyType:  pre.randomNextEnemyType(&mp.Resource), // TODO
		ThemeType:  pre.randomNextThemeType(&mp.Resource), // TODO
		Difficulty: pre.randomNextDifficulty(lucky),       // TODO
		Destiny:    pre.Destiny + destinyDiff,             // TODO
	}
	return quest
}

// GenerateQuest [pure]
func (mp *Map) GenerateQuest(lucky float64) {
	mp.GenerateNextQuest(lucky, 0, &Quest{
		Danger:    mp.Resource.Danger,
		QuestType: randomQuestType(),
	})
}
