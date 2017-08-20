package server

import (
)



// 处理连接数据
//func HandleConnection(conn *net.TCPConn) {
//	defer conn.Close()
//	fmt.Println(conn.RemoteAddr(), "new connection")
//	buf := make([]byte, 128)
//	for {
//		// 读取信息至缓冲区
//		length, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println(conn.RemoteAddr(), "connection closeed", err)
//			return
//		}
//
//		// 信息长度
//		msg, err := Decode(buf[:length])
//		if err != nil {
//			fmt.Println(err)
//		}
//		// TODO 处理接收到的消息
//		fmt.Println(conn.RemoteAddr(), "receive", msg)
//	}
//}
