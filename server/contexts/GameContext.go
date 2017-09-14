package contexts

import (
	"github.com/kataras/iris/websocket"
	"wandering-server/model"
)

type GameContext struct {
	Conn websocket.Connection `json:"-"`
	User model.User           `json:"user"`
	Log  func(...interface{})
	Emit func(string, interface{}) error
}

// 新建游戏内容实体
func NewGameContext(ctx *UserContext, game model.Game) GameContext {
	return GameContext{
		//todo
	}
}