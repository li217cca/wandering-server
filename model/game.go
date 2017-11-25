package model

import (
	"wandering-server/common"
)

/*
Game game struct
*/
type Game struct {
	ID       int    `json:"id,omitempty"`      // 游戏ID
	UserID   int    `json:"user_id,omitempty"` // 用户ID
	Name     string `json:"name,omitempty"`    // 名称
	BagID    int    `json:"bag_id,omitempty"`
	NowMapID int    `json:"map_id,omitempty"`
	Maps     []Map  `json:"maps,omitempty" gorm:"many2many:game_maps"`
}

/*
NewNativeGame New a native game
Type: not pure
UnitTest: false
*/
func NewNativeGame(UserID int, Name string) (game Game, err error) {
	// init map
	mp := NewMap(float64(common.NewGameGiftLucky+common.GetTodayLucky()), 0.)
	DB.Save(&mp)

	// init skill
	skills := []Skill{
		NewSkill("Body", SkillHitPointID, 0),
	}
	for index := range skills {
		DB.Save(&skills[index])
	}

	// init bag
	bag := NewBag()
	bag.calcCapacityWeight(skills)
	DB.Save(&bag)

	// init game
	game = Game{
		UserID:   UserID,
		Name:     Name,
		BagID:    bag.ID,
		NowMapID: mp.ID,
		Maps:     []Map{mp},
	}
	DB.Save(&game)

	// init char
	char := NewCharactor(game.ID, Name, CharNotInTeam)
	char.Skills = skills
	char.refreshHitPoint(skills)
	DB.Save(&char)
	// TODO: 更改各种方法为纯函数，编写单元测试
	return game, err
}
