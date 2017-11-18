package model

/*
Game game struct
*/
type Game struct {
	ID     int    `json:"id"`      // 游戏ID
	UserID int    `json:"user_id"` // 用户ID
	Name   string `json:"name"`    // 名称
	BagID  int    `json:"-"`       // 背包ID
}
