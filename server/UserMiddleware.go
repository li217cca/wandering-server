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
		fmt.Printf("(%s) [%s] %s: ", ctx.Conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"), user.Username)
		fmt.Println(args...)
	}
	return userContext{
		ctx.Conn,
		user,
		logger,
		//ctx.Conn.Emit,
	}, nil
}

func (ctx *userContext) OnSelect(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.GAME_SELECT, messageFunc)
}
func (ctx *userContext) OnCreate(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.GAME_CREATE, messageFunc)
}
func (ctx *userContext) sendGames() {
	var games []model.Game
	db.Model(&ctx.User).Related(&games)
	ctx.Conn.Emit(common.GAME_RECEIPT_LIST, games)
}
func (ctx *userContext) sendToken() error {
	token, err := model.NewUserToken(ctx.User, time.Now().Add(time.Second*300).Unix())
	if err != nil {
		return fmt.Errorf("\nuserContext.sendToken 01 \n%v", err)
	}
	return ctx.Conn.Emit(common.TOKEN_RECEIPT, token)
}
func (ctx *userContext) autoSendToken() {
	ticker := time.NewTicker(time.Minute * 3)
	for {
		if err := ctx.sendToken(); err != nil {
			ctx.Log("\nautoSendToken faild", err)
			return
		}
		<-ticker.C
	}
}

func handleUser(pctx *connContext, user model.User) error {
	ctx, err := newUserContext(pctx, user)
	if err != nil {
		return fmt.Errorf("\nhandleUser 01 %v", err)
	}
	ctx.Log("login")

	// auto send token
	go ctx.autoSendToken()

	// send game list..
	ctx.sendGames()

	ctx.OnSelect(func(gameID int) {
		ctx.Log("handUser OnSelect gameID=", gameID)
		game := model.Game{}
		if err := db.Model(model.Game{}).Where("id = ?", gameID).Find(&game).Error; err != nil {
			ctx.Log(fmt.Errorf("\nhandleUser ctx.OnSelect 01 \n%v", err))
		}
		db.Model(&game).Related(&game.Maps, "Maps")
		if err := handleGame(&ctx, game); err != nil {
			ctx.Log(fmt.Errorf("\nhandleUser ctx.OnSelect 02 \n%v", err))
		}
	})
	ctx.OnCreate(func(name string) {
		ctx.Log("handUser OnCreate name=", name)
		game, err := model.NewNativeGame(ctx.User.ID, name)
		if err != nil {
			ctx.Log(fmt.Errorf("\nhandleUser ctx.OnCreate 01 %v", err))
			return
		}
		if err := handleGame(&ctx, game); err != nil {
			ctx.Log(fmt.Errorf("\nhandleUser ctx.OnCreate 02 %v", err))
		}
		ctx.sendGames()
	})
	return nil
}
