package model

import (
	"wandering-server/common"
)

/*
Game game struct
*/
type Game struct {
	ID     int    `json:"id"`      // 游戏ID
	UserID int    `json:"user_id"` // 用户ID
	Name   string `json:"name"`    // 名称
	BagID  int    `json:"bag_id"`  // 背包ID
	MapID  int    `json:"map_id"`  // 地图ID
}

// commit not pure : save state to database
func (game *Game) commit() error {
	if DB.Where(Game{ID: game.ID}).RecordNotFound() {
		return DB.Model(game).Create(&game).Error
	}
	return DB.Where(Game{ID: game.ID}).Update(&game).Error
}

// func NewGame(UserID int, Name string, mp Map, bag Bag, Chars []Charactor) (game Game, err error) {
// }

// NewNativeGame ...
func NewNativeGame(UserID int, Name string) (game Game, err error) {
	mp := NewMap(common.NewGameGiftLucky+common.GetTodayLucky(), 0)
	bag := NewBag()

	game = Game{
		UserID: UserID,
		Name:   Name,
		MapID:  mp.ID,
		//BagID:  bag.ID,
	}
	game.commit()
	body := NewBodySkill("Body", 0)

	skills := []Skill{
		body,
	}
	char := NewCharactor(0, CharNotInTeam, skills)
	bag.calcCapacityWeight(char.Skills)
	// TODO: 更改各种方法为纯函数，编写单元测试
	return game, err
}
