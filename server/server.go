package server

import (

	"github.com/kataras/iris/websocket"
	"fmt"
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

	username := ""
	fmt.Println("connect from", conn.Context().RemoteAddr())
	// Read events from browser
	conn.On("login", func(name string) {
		fmt.Println(conn.Context().RemoteAddr(), "login:", username)
		conn.Emit("message", "login" + name)

		conn.On("battle", func(content string) {
			fmt.Println(conn.Context().RemoteAddr(), "on battle:", content)
		})
		conn.OnDisconnect(func() {
			fmt.Println(conn.Context().RemoteAddr(), "disconnect", username)
		})
	})
	conn.OnDisconnect(func() {
		fmt.Println(conn.Context().RemoteAddr(), "disconnect")
	})
}