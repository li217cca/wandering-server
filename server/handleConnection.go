package server

import (
	"github.com/kataras/iris/websocket"
	"wandering-server/server/contexts"
	"wandering-server/model"
)

func HandleConnection(conn websocket.Connection) {
	ctx := contexts.NewConnContext(conn, db)
	ctx.Log("connect")

	ctx.OnSignin(func(request interface{}) {
		reqmap := request.(map[string]interface{})
		username, ok1 := reqmap["username"].(string)
		password, ok2 := reqmap["password"].(string)
		password2, ok3 := reqmap["password2"].(string)
		name, ok4 := reqmap["name"].(string)
		if !(ok1 && ok2 && ok3 && ok4) {
			ctx.EmitError("输入格式不正确")
			return
		}
		if password != password2 {
			return
		}
		if len(username) < 6 {
			ctx.EmitError("用户名太短(<6)")
			return
		}
		if len(name) < 1 {
			ctx.EmitError("昵称")
			return
		}
		if len(password) < 6 {
			ctx.EmitError("密码太短(<6)")
			return
		}
		user := model.User{}
		if !ctx.DB.Model(model.User{}).Where("name = ?", name).First(&user).RecordNotFound() {
			return
		}
		if !ctx.DB.Model(model.User{}).Where("username = ?", username).First(&user).RecordNotFound() {
			ctx.EmitError("用户名已存在！")
			return
		}
		user = model.User{
			Username: username,
			Password: password,
			Name: name,
		}
		err := ctx.DB.Model(model.User{}).Create(&user).Error
		if err != nil {
			ctx.EmitError("未知错误" + err.Error())
			return
		}
		ctx.EmitSuccess("注册成功")
		HandleUser(&ctx, user)
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
		//go HandleUser(conn, user)
		HandleUser(&ctx, user)
	})

	ctx.OnDisconnect(func() {
		ctx.Log("disconnect")
	})
}
