package server

import (
	"fmt"
	"wandering-server/common"
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
	Info  model.Game
	Bag   model.Bag
	Map   model.Map
	Lucky int
}

// 新建游戏内容实体
func newGameContext(ctx *userContext, game model.Game) (gCtx gameContext, err error) {
	var bag model.Bag
	var mp model.Map
	db.Model(&game).Related(&bag)
	db.Model(&game).Related(&mp)
	gCtx = gameContext{
		Conn: ctx.Conn,
		User: ctx.User,
		Log: func(args ...interface{}) {
			var tmp []interface{}
			tmp = append(tmp, "["+game.Name+"]:")
			tmp = append(tmp, args...)
			ctx.Log(tmp...)
		},
		Emit: ctx.Conn.Emit,
		On:   ctx.Conn.On,
		Game: gameContainer{
			Info:  game,
			Bag:   bag,
			Map:   mp,
			Lucky: common.GetTodayLucky(), // FIXME: different lucky value in same day
		},
	}

	return gCtx, nil
}

func handleGame(pctx *userContext, game model.Game) error {
	ctx, err := newGameContext(pctx, game)
	if err != nil {
		return fmt.Errorf("\nhandleGame 01 %v", err)
	}
	ctx.Log("Game conn..")

	// send Game{}
	ctx.Emit(common.GAME_RECEIPT, game)

	// map search api
	ctx.On(common.MAP_SEARCH, func() {
		// TODO: map search
	})
	// TODO: 交互游戏信息
	return nil
}
