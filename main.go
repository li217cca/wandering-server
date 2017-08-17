package main

import (
	"wandering-server/model"
	"wandering-server/server"
)

func main() {
	conf := model.Config{
		8080,
	}
	ser := server.NewServer(conf)
	ser.Run()
}
