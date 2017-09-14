package server

import (
	"github.com/kataras/iris/websocket"
	"wandering-server/model"
	"wandering-server/server/handler"
	"wandering-server/common"
	"wandering-server/server/contexts"
)



func HandleUser(pctx *contexts.ConnContext, user model.User) {
	ctx := contexts.NewUserContext(pctx, user)
	ctx.Log("login")

	games := []model.Game{}
	ctx.DB.Model(&model.Game{}).Where("user_id = ?", ctx.User.ID).Find(&games)
	ctx.EmitGame(games)
	ctx.OnSelect(func(gameID int) {
		game := model.Game{}
		ctx.DB.Model(model.Game{}).Where("id = ?", gameID).Find(&game)
		HandleGame(&ctx, game)
	})
	ctx.OnCreate(func(name string) {
		if !ctx.DB.Model(model.Game{}).Where("name = ?", name).RecordNotFound() {
			ctx.EmitError("角色名已存在！")
			return
		}
		game := model.Game{
			UserID: ctx.User.ID,
			Name: name,
		}// todo 新建角色相关 数量上限
		ctx.DB.Model(model.Game{}).Create(&game)
		HandleGame(&ctx, game)
	})
}
