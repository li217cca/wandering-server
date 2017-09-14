package server

import (
	"wandering-server/server/contexts"
	"wandering-server/model"
	"wandering-server/common"
	"os/user"
)

func HandleGame(pctx *contexts.UserContext, game model.Game) {

	ctx := contexts.NewGameContext(pctx, game)

	// todo 交互游戏信息

	// put init state
	ctx.EmitParty()
	ctx.EmitUser()
	ctx.EmitBag()
	ctx.EmitScene()
	ctx.EmitTime()

	// register get request
	conn.On(common.GAME_GET_PARTY, ctx.EmitParty)
	conn.On(common.GAME_GET_TIME, ctx.EmitTime)
	conn.On(common.GAME_GET_SCENE, ctx.EmitScene)
	conn.On(common.GAME_GET_BAG, ctx.EmitBag)
	conn.On(common.GAME_GET_USER, ctx.EmitUser)

	// handle api
	conn.On(common.GAME_HANDLE_CARD, ctx.OnHandleCard)
	conn.On(common.GAME_HANDLE_ITEM, ctx.OnHandleItem)
	conn.On(common.GAME_HANDLE_TECH, ctx.OnHandleTech)

	conn.OnDisconnect(func() {
		ctx.Log(user.ID, " disconnect ...")
	})
}