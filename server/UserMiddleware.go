package server

import (
	"fmt"
	"time"
	"wandering-server/common"
	"wandering-server/model"

	"github.com/kataras/iris/websocket"
)

type userContext struct {
	Conn websocket.Connection `json:"-"`
	User model.User           `json:"user"`
	Log  func(...interface{})
	//Emit func(string, interface{}) error
}

// 新建用户内容实体
func newUserContext(ctx *connContext, user model.User) (userContext, error) {
	logger := func(args ...interface{}) {
		fmt.Printf("(%s) [%s]: ", ctx.Conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"))
		fmt.Println(args...)
	}
	return userContext{
		ctx.Conn,
		user,
		logger,
		//ctx.Conn.Emit,
	}, nil
}

func (ctx *userContext) EmitError(msg string) {
	ctx.Conn.Emit(common.GAME_ERROR, msg)
}
func (ctx *userContext) OnSelect(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.GAME_SELECT, messageFunc)
}
func (ctx *userContext) OnCreate(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.GAME_CREATE, messageFunc)
}

func handleUser(pctx *connContext, user model.User) error {
	ctx, err := newUserContext(pctx, user)
	if err != nil {
		return err
	}
	ctx.Log("login")

	games := []model.Game{}
	if err := db.Model(&model.Game{}).Where("user_id = ?", ctx.User.ID).Find(&games).Error; err != nil {
		return err
	}
	ctx.Conn.Emit(common.GAME_RECEIPT, games)
	ctx.OnSelect(func(gameID int) {
		game := model.Game{}
		if err := db.Model(model.Game{}).Where("id = ?", gameID).Find(&game).Error; err != nil {
			ctx.Log(err)
		}
		if err := handleGame(&ctx, game); err != nil {
			ctx.Log(err)
		}
	})
	ctx.OnCreate(func(name string) {
		if !db.Model(model.Game{}).Where("name = ?", name).RecordNotFound() {
			ctx.EmitError("角色名已存在！")
		}
		game := model.Game{
			UserID: ctx.User.ID,
			Name:   name,
		} // TODO: 新建角色相关 数量上限
		if err := db.Model(model.Game{}).Create(&game).Error; err != nil {
			ctx.Log(err)
		}
		if err := handleGame(&ctx, game); err != nil {
			ctx.Log(err)
		}
	})
	return nil
}
