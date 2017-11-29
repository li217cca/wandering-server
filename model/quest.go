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

	Destiny int   `json:"destiny,omitempty"` // 命运
	Length  int   `json:"length,omitempty"`
	PartyID int   `json:"party_id,omitempty"`
	Party   Party `json:"party,omitempty" gorm:"-"`
}

// AboutSize [pure]
func (quest *Quest) AboutSize() int {
	switch quest.QuestType {
	case QuestTypeBattleID:
		return 2
	case QuestTypeRaidID:
		return 6
	case QuestTypeMixedRaidID:
		return 8
	case QuestTypeDungeonID:
		return 23
	case QuestTypeMixedDungeonID:
		return 25
	case QuestTypeChaosDungeonID:
		return 15
	}
	return 1
}

// Quest const ...
const (
	// About CharType
	QuestDifficultyDefaultID = 0
	QuestDifficultyNormalID  = 100
	QuestDifficultyEliteID   = 200
	QuestDifficultyBossID    = 300
	QuestDifficultyPrizeID   = 400 // 奖励
	QuestDifficultyFriendID  = 500 // 友方

	// About Skill group
	EnemyThemeDefaultID  = 0   // 缺省主题
	EnemyThemeNormalID   = 100 // 普通
	EnemyThemeMagicID    = 200 // 魔法
	EnemyThemeScienceID  = 300 // 科学
	EnemyThemePunkID     = 400 // 朋克
	EnemyThemeDisasterID = 500 // 灾祸

	// About SpecialSkill, Skill group
	EnemyTypeDefaultID           = 0   // 缺省物种
	EnemyTypePlantID             = 100 // 植物
	EnemyTypePhytozoonID         = 200 // 素食动物
	EnemyTypeCarnivoreID         = 300 // 肉食动物
	EnemyTypeCivilizationID      = 400 // 文明
	EnemyTypeCivilizationExileID = 410 // 放逐文明
	EnemyTypeLegendID            = 500 // 传说

	/**
	 * AboutSize
	 * randomQuestLength
	 * randomQuestType
	 * isMixed
	 * QuestTypeToString
	 */
	// About Battle size, Difficulty, CharType weight
	QuestTypeDefaultID        = 0
	QuestTypeBattleID         = 200 // [1,   2] 一种种族，概率Prize
	QuestTypeRaidID           = 300 // [3,   9] 一种种族，概率Prize，概率BOSS
	QuestTypeMixedRaidID      = 310 // [4,  12] 几种种族，概率Prize，概率BOSS
	QuestTypeDungeonID        = 400 // [9,  36] 一种种族，少量Prize，少量BOSS
	QuestTypeMixedDungeonID   = 410 // [10, 40] 几种种族，少量Prize，少量BOSS
	QuestTypeChaosDungeonID   = 430 // [5,  25] 几种Elite/BOSS，Prize
	QuestTypeAncientDungeonID = 450 // [X,   X] 远古地牢 // TODO
)

func randomQuestLength(questType int) int {
	switch questType {
	case QuestTypeBattleID:
		return common.Float(10, 19) / 10
	case QuestTypeRaidID:
		return int(common.GaussianlRandF3(6, 3))
	case QuestTypeMixedRaidID:
		return int(common.GaussianlRandF3(8, 4))
	case QuestTypeDungeonID:
		return int(common.GaussianlRandF3(23, 15))
	case QuestTypeMixedDungeonID:
		return int(common.GaussianlRandF3(25, 16))
	case QuestTypeChaosDungeonID:
		return int(common.GaussianlRandF3(15, 10))
	case QuestTypeAncientDungeonID:
		return 0
	}
	return 1
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
		{10000, QuestTypeBattleID},
		{int(500 * landAreaAddition), QuestTypeRaidID},
		{int(150 * landAreaAddition), QuestTypeMixedRaidID},
		{int(100 * landAreaAddition * luckyAddition), QuestTypeDungeonID},
		{int(30 * landAreaAddition * luckyAddition), QuestTypeMixedDungeonID},
		{int(20 * landAreaAddition * luckyAddition), QuestTypeChaosDungeonID},
		{0, QuestTypeAncientDungeonID},
	}
	get, ok := rou.Get().(int)
	if ok {
		return get
	}
	return QuestTypeDefaultID
}

func (quest *Quest) isMixed() bool {
	switch quest.QuestType {
	case QuestTypeMixedRaidID:
		return true
	case QuestTypeMixedDungeonID:
		return true
	case QuestTypeChaosDungeonID:
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
	disasterAddition := math.Sqrt(res.DisasterResource / 100)
	rou := common.Roulette{
		{int(100 * math.Sqrt(res.PlantResource/10000*disasterAddition)), EnemyTypePlantID},
		{int(100 * math.Sqrt(res.PhytozoonResource/1000)), EnemyTypePhytozoonID},
		{int(100 * math.Sqrt(res.CarnivoreResource/200)), EnemyTypeCarnivoreID},
		{int(100 * math.Sqrt(res.CivilizationResource/200)), EnemyTypeCivilizationID},
		{int(100 * math.Sqrt(res.CivilizationResource/1000)), EnemyTypeCivilizationExileID},
		{int(100 * math.Sqrt(res.LegendResource/1000)), EnemyTypeLegendID},
		{int(notMix), quest.EnemyType},
	}
	get, ok := rou.Get().(int)
	if ok {
		return get
	}
	return EnemyTypeDefaultID
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
	get, ok := rou.Get().(int)
	if ok {
		return get
	}
	return EnemyThemeDefaultID
}

// randomNextDifficulty [pure]
func (quest *Quest) randomNextDifficulty(lucky float64) int {
	destinyAddition := math.Sqrt(float64(quest.Destiny / 600))
	luckyAddition := math.Sqrt(lucky / 50)
	sizeAddition := math.Sqrt(float64(quest.AboutSize()) / 7)
	lengthAddition := 1 / math.Sqrt(float64(quest.Length+1))
	rou := common.Roulette{
		{int(100), QuestDifficultyNormalID},
		{int(5 + 50*destinyAddition), QuestDifficultyEliteID},
		{int(100 * destinyAddition * sizeAddition * lengthAddition), QuestDifficultyBossID},
		{int(5 + 50*destinyAddition*luckyAddition*lengthAddition), QuestDifficultyPrizeID},
		{int(10 * luckyAddition * sizeAddition * lengthAddition), QuestDifficultyFriendID},
	}
	get, ok := rou.Get().(int)
	if ok {
		return get
	}
	return QuestDifficultyDefaultID
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
	case QuestDifficultyPrizeID:
		destiny *= 0.1
	case QuestDifficultyFriendID:
		destiny *= 0.1
	case QuestDifficultyEliteID:
		destiny *= 2
	case QuestDifficultyBossID:
		destiny *= 5
	}
	switch quest.EnemyType {
	case EnemyTypeLegendID:
		destiny *= 2
	case EnemyTypeCivilizationExileID:
		destiny *= 1.3
	case EnemyTypeCivilizationID:
		destiny *= 1.3
	case EnemyTypePlantID:
		destiny *= 0.7
	}
	return int(destiny)
}

// UseDistiny [pure]
func (quest *Quest) UseDistiny(getDestiny int) int {
	destiny := float64(quest.Destiny + getDestiny)
	switch quest.Difficulty {
	case QuestDifficultyFriendID:
		r := common.GetRand()
		return int(common.FloatF(0, destiny*r.Float64()*r.Float64()*r.Float64()))
	case QuestDifficultyBossID:
		return int(common.FloatF(destiny*0.99, destiny))
	case QuestDifficultyPrizeID:
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
	case QuestTypeBattleID:
		return "战斗"
	case QuestTypeRaidID:
		return "突袭"
	case QuestTypeMixedRaidID:
		return "混合突袭"
	case QuestTypeDungeonID:
		return "地牢"
	case QuestTypeMixedDungeonID:
		return "混合地牢"
	case QuestTypeChaosDungeonID:
		return "混沌地牢"
	case QuestTypeAncientDungeonID:
		return "远古地牢"
	}
	return "未知Quest"
}

// EnemyTypeToString [pure]
func (quest *Quest) EnemyTypeToString() string {
	switch quest.EnemyType {
	case EnemyTypePlantID:
		return "植物"
	case EnemyTypePhytozoonID:
		return "素食动物"
	case EnemyTypeCarnivoreID:
		return "肉食动物"
	case EnemyTypeCivilizationID:
		return "文明"
	case EnemyTypeCivilizationExileID:
		return "放逐文明"
	case EnemyTypeLegendID:
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
	case QuestDifficultyNormalID:
		return ""
	case QuestDifficultyEliteID:
		return "精英"
	case QuestDifficultyBossID:
		return "BOSS"
	case QuestDifficultyFriendID:
		return "友好"
	case QuestDifficultyPrizeID:
		return "奖励"
	}
	return "未知难度"
}

func (quest *Quest) getName() string {
	name := quest.DifficultyToString() + quest.QuestTypeToString()
	if !(quest.Difficulty == QuestDifficultyFriendID || quest.Difficulty == QuestDifficultyPrizeID) {
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
		%sQuest[%s]{ID: %d, PreID: %d, MapID: %d, Danger: %.1f, Destiny: %d, Length: %d}
		%s	%s`,
		diff, quest.getName(), quest.ID, quest.PreID, quest.MapID, quest.Danger, quest.Destiny, quest.Length, diff, quest.Party.ToString(),
	)
	return str
}
