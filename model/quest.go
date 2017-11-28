package model

import (
	"fmt"
	"math"
	"wandering-server/common"
)

/*
Quest ...
*/
type Quest struct {
	ID     int     `json:"id,omitempty"`
	Key    string  `json:"key,omitempty" gorm:"-"`
	PreID  int     `json:"pre_id,omitempty"` // 上一个ID
	MapID  int     `json:"map_id,omitempty" gorm:"not null"`
	Danger float64 `json:"danger,omitempty" gorm:"not null"`

	// Name string `json:"name,omitempty"    gorm:"not null"`
	QuestType  int `gorm:"not null" json:"quest_type,omitempty"` // 种类...
	EnemyType  int `gorm:"not null" json:"enemy_type,omitempty"` // 植物/肉/素食动物/文明/传说...
	ThemeType  int `gorm:"not null" json:"theme_type,omitempty"` // 普通/科学/魔法/朋克/灾祸...
	Difficulty int `gorm:"not null" json:"difficulty,omitempty"` // 普通/精英/BOSS

	Destiny int `json:"destiny,omitempty"` // 命运
	Length  int `json:"length,omitempty"`
}

// AboutSize [pure]
func (quest *Quest) AboutSize() int {
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

func randomQuestLength(questType int) int {
	switch questType {
	case QuestBattleID:
		return common.Float(10, 19) / 10
	case QuestRaidID:
		return int(common.GaussianlRandF3(6, 3))
	case QuestMixedRaidID:
		return int(common.GaussianlRandF3(8, 4))
	case QuestDungeonID:
		return int(common.GaussianlRandF3(23, 15))
	case QuestMixedDungeonID:
		return int(common.GaussianlRandF3(25, 16))
	case QuestChaosDungeonID:
		return int(common.GaussianlRandF3(15, 10))
	case QuestAncientDungeonID:
		return 0
	}
	return 0
}

// randomQuestType [pure]
func randomQuestType(lucky float64, landArea float64) int {
	if lucky < 0 {
		lucky = 0
	}
	if landArea < 0 {
		landArea = 0
	}
	landAreaAddition := math.Sqrt(landArea / 1500)
	luckyAddition := math.Sqrt(lucky / 30)
	rou := common.Roulette{
		{10000, QuestBattleID},
		{int(500 * landAreaAddition), QuestRaidID},
		{int(150 * landAreaAddition), QuestMixedRaidID},
		{int(100 * landAreaAddition * luckyAddition), QuestDungeonID},
		{int(30 * landAreaAddition * luckyAddition), QuestMixedDungeonID},
		{int(20 * landAreaAddition * luckyAddition), QuestChaosDungeonID},
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
		{int(100 * math.Sqrt(res.PlantResource/10000)), EnemyPlantID},
		{int(100 * math.Sqrt(res.PhytozoonResource/1000)), EnemyPhytozoonID},
		{int(100 * math.Sqrt(res.CarnivoreResource/200)), EnemyCarnivoreID},
		{int(100 * math.Sqrt(res.CivilizationResource/200)), EnemyCivilizationID},
		{int(100 * math.Sqrt(res.CivilizationResource/1000)), EnemyCivilizationExileID},
		{int(100 * math.Sqrt(res.LegendResource/1000)), EnemyLegengID},
		{int(notMix), quest.EnemyType},
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
		{int(100 * math.Sqrt(res.Area()/1000)), EnemyThemeNormalID},
		{int(100 * math.Sqrt(res.MagicResource/1000)), EnemyThemeMagicID},
		{int(100 * math.Sqrt(res.ScienceResource/1000)), EnemyThemeScienceID},
		{int(100 * math.Sqrt(res.MagicResource/500) * math.Sqrt(res.ScienceResource/500)), EnemyThemePunkID},
		{int(100 * math.Sqrt(res.DisasterResource/1000)), EnemyThemeDisasterID},
		{int(notMix), quest.ThemeType},
	}
	return rou.Get().(int)
}

// randomNextDifficulty [pure]
func (quest *Quest) randomNextDifficulty(lucky float64) int {
	destinyAddition := math.Sqrt(float64(quest.Destiny / 600))
	luckyAddition := math.Sqrt(lucky / 50)
	sizeAddition := math.Sqrt(float64(quest.AboutSize()) / 7)
	lengthAddition := 1 / math.Sqrt(float64(quest.Length+1))
	rou := common.Roulette{
		{int(100), EnemyDifficultyNormalID},
		{int(5 + 50*destinyAddition), EnemyDifficultyEliteID},
		{int(100 * destinyAddition * sizeAddition * lengthAddition), EnemyDifficultyBossID},
		{int(5 + 50*destinyAddition*luckyAddition*lengthAddition), EnemyDifficultyPrizeID},
		{int(10 * luckyAddition * sizeAddition * lengthAddition), EnemyDifficultyFriendID},
	}
	return rou.Get().(int)
}

// GetDistiny [pure]
func (quest *Quest) GetDistiny() int {
	destiny := common.GaussianlRandF3(200, 50)
	switch quest.ThemeType {
	case EnemyThemeDisasterID:
		destiny *= 1.2
	case EnemyThemePunkID:
		destiny *= 1.1
	case EnemyThemeNormalID:
		destiny *= 0.9
	}
	switch quest.Difficulty {
	case EnemyDifficultyPrizeID:
		destiny *= 0.1
	case EnemyDifficultyFriendID:
		destiny *= 0.1
	case EnemyDifficultyEliteID:
		destiny *= 2
	case EnemyDifficultyBossID:
		destiny *= 5
	}
	switch quest.EnemyType {
	case EnemyLegengID:
		destiny *= 2
	case EnemyCivilizationExileID:
		destiny *= 1.3
	case EnemyCivilizationID:
		destiny *= 1.3
	case EnemyPlantID:
		destiny *= 0.7
	}
	return int(destiny)
}

// UseDistiny [pure]
func (quest *Quest) UseDistiny(getDestiny int) int {
	destiny := float64(quest.Destiny + getDestiny)
	switch quest.Difficulty {
	case EnemyDifficultyFriendID:
		r := common.GetRand()
		return int(common.FloatF(0, destiny*r.Float64()*r.Float64()*r.Float64()))
	case EnemyDifficultyBossID:
		return int(common.FloatF(destiny*0.99, destiny))
	case EnemyDifficultyPrizeID:
		return int(common.FloatF(destiny*0.1, destiny*0.8))
	}
	return int(common.FloatF(0, destiny*0.05))
}

// IsEnd [pure]
func (quest *Quest) IsEnd(destinyDiff int) bool {
	if quest.AboutSize() > quest.Length*2 {
		if math.Sqrt(float64(quest.Destiny+destinyDiff)/100) < common.FloatF(0.5, 1.3) {
			return true
		}
	}
	return false
}

// GenerateNextQuest [pure]
func (mp *Map) GenerateNextQuest(lucky float64, destinyDiff int, pre *Quest) Quest {
	quest := Quest{
		Key:        common.GenerateKey(12),
		PreID:      pre.ID,
		Danger:     common.FloatF(pre.Danger-1, pre.Danger+2),
		QuestType:  pre.QuestType,
		EnemyType:  pre.randomNextEnemyType(&mp.Resource),
		ThemeType:  pre.randomNextThemeType(&mp.Resource),
		Difficulty: pre.randomNextDifficulty(lucky),
		Destiny:    pre.Destiny + destinyDiff,
		Length:     pre.Length - 1,
	}
	return quest
}

// GenerateQuest [pure]
func (mp *Map) GenerateQuest(lucky float64) Quest {
	questType := randomQuestType(lucky, mp.Resource.Area())
	quest := mp.GenerateNextQuest(lucky, 0, &Quest{
		Danger:    mp.Resource.Danger,
		QuestType: questType,
		Length:    randomQuestLength(questType) + 1,
	})
	return quest
}

// QuestTypeToString [pure]
func (quest *Quest) QuestTypeToString() string {
	switch quest.QuestType {
	case QuestBattleID:
		return "战斗"
	case QuestRaidID:
		return "突袭"
	case QuestMixedRaidID:
		return "混合突袭"
	case QuestDungeonID:
		return "地牢"
	case QuestMixedDungeonID:
		return "混合地牢"
	case QuestChaosDungeonID:
		return "混沌地牢"
	case QuestAncientDungeonID:
		return "远古地牢"
	}
	return "未知Quest"
}

// EnemyTypeToString [pure]
func (quest *Quest) EnemyTypeToString() string {
	switch quest.EnemyType {
	case EnemyPlantID:
		return "植物"
	case EnemyPhytozoonID:
		return "素食动物"
	case EnemyCarnivoreID:
		return "肉食动物"
	case EnemyCivilizationID:
		return "文明"
	case EnemyCivilizationExileID:
		return "放逐文明"
	case EnemyLegengID:
		return "传说"
	}
	return "未知种族"
}

// ThemeTypeToString [pure]
func (quest *Quest) ThemeTypeToString() string {
	switch quest.ThemeType {
	case EnemyThemeNormalID:
		return ""
	case EnemyThemeMagicID:
		return "魔法"
	case EnemyThemeScienceID:
		return "科技"
	case EnemyThemePunkID:
		return "朋克"
	case EnemyThemeDisasterID:
		return "灾祸"
	}
	return "未知类型"
}

// DifficultyToString [pure]
func (quest *Quest) DifficultyToString() string {
	switch quest.Difficulty {
	case EnemyDifficultyNormalID:
		return ""
	case EnemyDifficultyEliteID:
		return "精英"
	case EnemyDifficultyBossID:
		return "BOSS"
	case EnemyDifficultyFriendID:
		return "友好"
	case EnemyDifficultyPrizeID:
		return "奖励"
	}
	return "未知难度"
}

func (quest *Quest) getName() string {
	name := quest.DifficultyToString() + quest.QuestTypeToString()
	if !(quest.Difficulty == EnemyDifficultyFriendID || quest.Difficulty == EnemyDifficultyPrizeID) {
		name = quest.ThemeTypeToString() +
			quest.EnemyTypeToString() + name
	}
	return name
}

// ToString [pure]
func (quest *Quest) ToString() string {
	diff := ``
	if quest.PreID != 0 {
		diff = `	`
	}
	str := fmt.Sprintf(`
		%sQuest[%s]{ID: %d, PreID: %d, MapID: %d, Danger: %.1f, Destiny: %d, Length: %d}`,
		diff, quest.getName(), quest.ID, quest.PreID, quest.MapID, quest.Danger, quest.Destiny, quest.Length,
	)
	return str
}
