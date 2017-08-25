package server

import (
	"fmt"
	"github.com/kataras/iris/websocket"
	"time"
	"wandering-server/model"
)

type DataPack struct {
	Chars []struct {
		model.Character
		Skills []model.Skill
	}
}

func HandleUser(conn websocket.Connection, user model.User) {
	log := func(args ...interface{}) {
		fmt.Printf("(%s) [%s] %s: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"), user.ID)
		fmt.Println(args...)
	}
	log("login")
	data := DataPack{}
	db.Model(&model.Character{}).Where("user_id = ?", user.ID).Find(&data.Chars)

	conn.OnDisconnect(func() {
		log(user.ID, " disconnect ...")
	})
}
