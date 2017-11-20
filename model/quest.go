package model

/*
Quest ...
*/
type Quest struct {
	ID     int
	NextID int
	MapID  int    `gorm:"not null"`
	Name   string `gorm:"not null"`
	Type   int    `gorm:"not null"`
	Level  int    `gorm:"not null"`
}

/*
Quest const ...
*/
const (
	QuestNormalBattleID     = 100
	QuestNormalBossBattleID = 101
	QuestGetWeaponID        = 202
	// TODO: more quest type
)

/*
Map.generateQuest Generate quest from lucky, danger, miracle ...
Type: pure
UnitTest: false
*/
func (mp *Map) generateQuest(lucky int) {
	// TODO: ...
	// quest := Quest{
	// }
}
