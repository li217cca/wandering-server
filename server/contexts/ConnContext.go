package contexts

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/websocket"
	"wandering-server/common"
)

type ConnContext struct {
	Conn websocket.Connection `json:"-"`
	Log  func(...interface{})
	DB   *gorm.DB
}

// 新建连接内容实体
func NewConnContext(conn websocket.Connection, db *gorm.DB) ConnContext {
	logger := func(args ...interface{}) {
		fmt.Printf("(%s) [%s]: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"))
		fmt.Println(args...)
	}
	ctx := ConnContext{
		conn,      // socket 连接
		logger,    // 日志器
		db,        // 数据库
	}
	return ctx
}

func (ctx * ConnContext) EmitError(msg string)  {
	ctx.Conn.Emit(common.AUTH_ERROR, msg)
}
func (ctx * ConnContext) EmitSuccess(msg string)  {
	ctx.Conn.Emit(common.AUTH_SUCCESS, msg)
}
func (ctx * ConnContext) OnLogin(messageFunc websocket.MessageFunc)  {
	ctx.Conn.On(common.AUTH_LOGIN, messageFunc)
}
func (ctx * ConnContext) OnSignin(messageFunc websocket.MessageFunc)  {
	ctx.Conn.On(common.AUTH_SIGNIN, messageFunc)
}
func (ctx * ConnContext) OnDisconnect(disconnectFunc websocket.DisconnectFunc)  {
	ctx.Conn.OnDisconnect(disconnectFunc)
}