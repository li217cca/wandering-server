package server

import (
	"wandering-server/model"

	"github.com/kataras/iris/websocket"
)

type gameContext struct {
	Conn websocket.Connection `json:"-"`
	User model.User           `json:"user"`
	Log  func(...interface{})
	Emit func(string, interface{}) error
	Game gameContainer
}
type gameContainer struct {
	Info model.Game
	Bag  model.Bag
}

// 新建游戏内容实体
func newGameContext(ctx *userContext, game model.Game) (gameContext, error) {
	bag, err := model.GetBagByID(game.BagID)
	if err != nil {
		return gameContext{}, err
	}
	return gameContext{
		Conn: ctx.Conn,
		User: ctx.User,
		Log:  ctx.Log,
		Emit: ctx.Conn.Emit,
		Game: gameContainer{
			game,
			bag,
		},
	}, nil
}

func handleGame(pctx *userContext, game model.Game) error {
	ctx, err := newGameContext(pctx, game)
	if err != nil {
		return err
	}

	ctx.Log("GAME_CONN")
	// todo 交互游戏信息
	return nil
}
