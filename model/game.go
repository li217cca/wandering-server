package model

import (
	"wandering-server/common"
)

/*
Game game struct
*/
type Game struct {
	ID       int         `json:"id,omitempty"`      // 游戏ID
	UserID   int         `json:"user_id,omitempty"` // 用户ID
	Name     string      `json:"name,omitempty"`    // 名称
	BagID    int         `json:"bag_id,omitempty"`
	NowMapID int         `json:"map_id,omitempty"`
	Maps     []Map       `json:"maps,omitempty" gorm:"many2many:game_maps"`
	Chars    []Charactor `json:"chars,omitempty" gorm:"many2many:game_chars"`
	Partys   []Party     `json:"partys,omitempty"`
}

/*
NewNativeGame New a native game
Type: not pure
UnitTest: false
*/
func NewNativeGame(UserID int, Name string) (game Game, err error) {

	// init char
	char := NewCharactor(Name, []Skill{
		NewSkill("Body", SkillHitPointBaseID, 0),
	})
	DB.Save(&char)

	// init bag
	bag := NewBag()
	bag.calcCapacityWeight(char.Skills)
	DB.Save(&bag)

	// init game
	game = Game{
		UserID: UserID,
		Name:   Name,
		BagID:  bag.ID,
	}
	DB.Save(&game)

	// init map
	mp := NewMap(float64(common.NewGameGiftLucky+common.GetTodayLucky(game.ID)), 0)

	// assign map to game
	game.NowMapID = mp.ID
	game.Maps = []Map{mp}

	party := Party{
		Chars: [6]Charactor{char},
	}
	DB.Save(&party)
	game.Partys = append(game.Partys, party)
	game.Chars = append(game.Chars, char)
	DB.Save(&game)
	// TODO: 更改各种方法为纯函数，编写单元测试
	return game, err
}
