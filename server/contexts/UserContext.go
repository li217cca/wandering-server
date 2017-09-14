package contexts

import (
	"github.com/kataras/iris/websocket"
	"wandering-server/model"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"wandering-server/common"
)

type UserContext struct {
	Conn websocket.Connection `json:"-"`
	User model.User           `json:"user"`
	DB *gorm.DB
	Log  func(...interface{})
	//Emit func(string, interface{}) error
}

// 新建用户内容实体
func NewUserContext(ctx *ConnContext, user model.User) UserContext {
	logger := func(args ...interface{}) {
		fmt.Printf("(%s) [%s] %s: ", ctx.Conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"), user.Name)
		fmt.Println(args...)
	}
	return UserContext{
		ctx.Conn,
		user,
		ctx.DB,
		logger,
		//ctx.Conn.Emit,
	}
}

func (ctx *UserContext) EmitGame(games []model.Game)  {
	ctx.Conn.Emit(common.GAME_RECEIPT, games)
}
func (ctx *UserContext) EmitError(msg string)  {
	ctx.Conn.Emit(common.GAME_ERROR, msg)
}
func (ctx *UserContext) OnSelect(messageFunc websocket.MessageFunc)  {
	ctx.Conn.On(common.GAME_SELECT, messageFunc)
}
func (ctx *UserContext) OnCreate(messageFunc websocket.MessageFunc)  {
	ctx.Conn.On(common.GAME_CREATE, messageFunc)
}