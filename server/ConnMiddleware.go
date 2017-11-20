package server

import (
	"fmt"
	"time"

	"wandering-server/common"
	"wandering-server/model"

	"github.com/kataras/iris/websocket"
)

type connContext struct {
	Conn websocket.Connection `json:"-"`
	Log  func(...interface{})
}

// 新建连接内容实体
func newConnContext(conn websocket.Connection) connContext {
	logger := func(args ...interface{}) {
		fmt.Printf("(%s) [%s]: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"))
		fmt.Println(args...)
	}
	ctx := connContext{
		conn,   // socket 连接
		logger, // 日志器
	}
	return ctx
}

func (ctx *connContext) EmitError(msg string) {
	ctx.Conn.Emit(common.AUTH_ERROR, msg)
}
func (ctx *connContext) EmitSuccess(msg string) {
	ctx.Conn.Emit(common.AUTH_SUCCESS, msg)
}
func (ctx *connContext) OnLogin(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.AUTH_LOGIN, messageFunc)
}
func (ctx *connContext) OnSignin(messageFunc websocket.MessageFunc) {
	ctx.Conn.On(common.AUTH_SIGNIN, messageFunc)
}

// HandleConnection ...
func HandleConnection(conn websocket.Connection) {
	ctx := newConnContext(conn)
	ctx.Log("connect")

	ctx.OnSignin(func(request interface{}) {
		reqmap := request.(map[string]interface{})
		username, ok1 := reqmap["username"].(string)
		password, ok2 := reqmap["password"].(string)
		if !(ok1 && ok2) {
			ctx.EmitError("输入格式不正确")
			return
		}
		if len(username) < 6 {
			ctx.EmitError("用户名太短(<6)")
			return
		}
		if len(password) < 6 {
			ctx.EmitError("密码太短(<6)")
			return
		}
		user := model.User{}
		if !db.Model(model.User{}).Where("username = ?", username).First(&user).RecordNotFound() {
			ctx.EmitError("用户名已存在！")
			return
		}
		user = model.User{
			Username: username,
			Password: password,
		}
		if err := db.Model(model.User{}).Create(&user).Error; err != nil {
			ctx.EmitError("未知错误" + err.Error())
			ctx.Log(fmt.Errorf("\n HandleConnection 01 \n%v", err))
			return
		}
		ctx.EmitSuccess("注册成功")
		if err := handleUser(&ctx, user); err != nil {
			ctx.Log(fmt.Errorf("\n HandleConnection 02 \n%v", err))
		}
	})
	ctx.OnLogin(func(request interface{}) {
		username, ok1 := request.(map[string]interface{})["username"].(string)
		password, ok2 := request.(map[string]interface{})["password"].(string)
		if !(ok1 && ok2) {
			ctx.EmitError("输入格式不正确")
			return
		}
		user, err := model.GetUserByAuthenticate(username, password)
		if err != nil {
			ctx.EmitError("用户名或密码错误")
			return
		}
		ctx.EmitSuccess("登陆成功")
		if err := handleUser(&ctx, user); err != nil {
			ctx.Log(err)
		}
	})

	ctx.Conn.OnDisconnect(func() {
		ctx.Log("disconnect")
	})
}
