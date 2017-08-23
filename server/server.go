package server

import (

	"github.com/kataras/iris/websocket"
	"fmt"
	"time"
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
func HandleConnection(conn websocket.Connection)  {

	log := func(text ...interface{}) {
		fmt.Printf("[%s] (%s): %s\n", time.Now().Format("01/02 15:04:05.0"), conn.Context().RemoteAddr(), text)
	}
	username := ""
	log("connect")
	// Read events from browser
	conn.OnMessage(func(bytes []byte) {
		log("message:", string(bytes))
	})
	conn.On("login", func(request interface{}) {
		//request.(map[string]interface{})["username"]
		//request.(map[string]interface{})["password"]
		log("login")
		conn.Emit("login", "success")

		conn.On("battle", func(content string) {
			fmt.Println(conn.Context().RemoteAddr(), "on battle:", content)
		})
		conn.OnDisconnect(func() {
			fmt.Println(conn.Context().RemoteAddr(), "disconnect", username)
		})
	})
	conn.OnDisconnect(func() {
		log("disconnect")
	})
}