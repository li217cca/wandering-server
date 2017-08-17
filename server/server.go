package server

import (
	"fmt"
	"net"
	"wandering-server/model"
)

type Server struct {
	Addr     net.TCPAddr
	listener net.TCPListener
}

func NewServer(conf model.Config) *Server {
	addr := net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: conf.Port}
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Run() {
	// 注册TCP服务在指定端口
	ln, err := net.ListenTCP("tcp", &s.Addr)
	defer ln.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("listen in", s.Addr.String(), "...\n")

	// 接受连接
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go HandleConnection(conn)
	}
}
