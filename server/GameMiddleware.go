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
	Lucky float64
}

// 新建游戏内容实体
func newGameContext(ctx *userContext, game model.Game) (gCtx gameContext, err error) {
	var bag model.Bag
	mp := model.Map{}
	// db.Model(&game).Related(&bag)
	for index := range game.Maps {
		if game.NowMapID == game.Maps[index].ID {
			mp = game.Maps[index]
		}
	}
	if mp.ID == 0 {
		mp = game.Maps[0]
		game.NowMapID = mp.ID
		db.Save(&game)
	}

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
			Lucky: common.GetTodayLucky(game.ID), // FIXME: different lucky value in same day
		},
	}

	return gCtx, nil
}

func handleGame(pctx *userContext, game model.Game) error {
	ctx, err := newGameContext(pctx, game)
	if err != nil {
		return fmt.Errorf("\nhandleGame 01 %v", err)
	}
	ctx.Log("Join Game")

	gameContainers[ctx.Game.Info.ID] = &ctx.Game
	// send Game{}
	ctx.Emit(common.GAME_RECEIPT, ctx.Game.Info)
	ctx.Emit(common.MAP_RECEIPT, ctx.Game.Map)
	ctx.Emit(common.BAG_RECEIPT, ctx.Game.Bag)

	// map search api
	ctx.On(common.MAP_SEARCH, func() {
		ctx.Game.Map.Search(ctx.Game.Lucky)
		ctx.Emit(common.MAP_RECEIPT, ctx.Game.Map)
	})
	// TODO: 交互游戏信息
	ctx.Conn.OnDisconnect(func() {
		ctx.Log("Leave Game")
		delete(gameContainers, ctx.Game.Info.ID)
	})
	return nil
}
