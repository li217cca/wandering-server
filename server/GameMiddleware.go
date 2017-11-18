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
	On   func(string, websocket.MessageFunc)
}
type gameContainer struct {
	Info model.Game
	Bag  model.Bag
	Map  model.Map
}

// 新建游戏内容实体
func newGameContext(ctx *userContext, game model.Game) (gCtx gameContext, err error) {
	gCtx = gameContext{
		Conn: ctx.Conn,
		User: ctx.User,
		Log:  ctx.Log,
		Emit: ctx.Conn.Emit,
		On:   ctx.Conn.On,
		Game: gameContainer{
			Info: game,
		},
	}
	if gCtx.Game.Bag, err = model.GetBagByID(game.BagID); err != nil {
		return gCtx, err
	}
	if gCtx.Game.Map, err = model.GetMapByID(game.MapID); err != nil {
		return gCtx, err
	}
	return gCtx, err
}

func handleGame(pctx *userContext, game model.Game) error {
	ctx, err := newGameContext(pctx, game)
	if err != nil {
		return err
	}

	ctx.Log("GAME_CONN")
	// TODO: 交互游戏信息
	return nil
}
