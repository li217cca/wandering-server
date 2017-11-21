package model

/*
Quest ...
*/
type Quest struct {
	ID     int    `json:"id,omitempty"`
	NextID int    `json:"next_id,omitempty"`
	MapID  int    `json:"map_id,omitempty"  gorm:"not null" `
	Name   string `json:"name,omitempty"    gorm:"not null"`
	Type   int    `json:"type,omitempty"    gorm:"not null"`
	Level  int    `json:"level,omitempty"   gorm:"not null"`
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
