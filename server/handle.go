package server

import (
	"fmt"
	"net"
)



func HandleConnection(conn *net.TCPConn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr(), "new connection")
	buf := make([]byte, 128)
	for {
		// 读取信息至缓冲区
		length, err := conn.Read(buf)
		if err != nil {
			fmt.Println("a connection closeed in", conn.RemoteAddr(), err)
			return
		}

		// 信息长度
		msg, err := Decode(buf[:length])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn.RemoteAddr(), "receive", msg)
	}
}
