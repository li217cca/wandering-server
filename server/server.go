package server

import (
	"fmt"
	"github.com/kataras/iris/websocket"
	"time"
	"wandering-server/model"
)

//
//// 运行服务
//func (s *Server) Run() {
//	// 注册TCP服务在指定端口
//	ln, err := net.ListenTCP("tcp", &s.Addr)
//	defer ln.Close()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("listen in", s.Addr.String(), "...\n")
//
//	// 接受连接
//	for {
//		conn, err := ln.AcceptTCP()
//		if err != nil {
//			panic(err)
//		}
//		go HandleConnection(conn)
//	}
//}

func HandleConnection(conn websocket.Connection) {
	log := func(args ...interface{}) {
		fmt.Printf("(%s) [%s]: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"))
		fmt.Println(args...)
	}

	//log("connect")
	// Auth api
	conn.On("AUTH_LOGIN", func(request interface{}) {
		username, ok1 := request.(map[string]interface{})["username"].(string)
		password, ok2 := request.(map[string]interface{})["password"].(string)
		if !(ok1 && ok2) {
			conn.Emit("AUTH_ERROR", "输入格式不正确")
			return
		}
		user := model.User{}
		err := db.Model(&model.User{}).Where(model.User{
			Username: username,
			Password: password,
		}).First(&user).Error
		if err != nil {
			conn.Emit("AUTH_ERROR", "用户名或密码错误")
			return
		}
		conn.Emit("AUTH_SUCCESS", "登陆成功")
		//go HandleUser(conn, user)
		HandleUser(conn, user)
	})

	conn.OnDisconnect(func() {
		//log("disconnect")
	})
}
