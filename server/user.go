package server

import (
	"fmt"
	"github.com/kataras/iris/websocket"
	"time"
	"wandering-server/model"
)

type DataPack struct {
	Characters []struct {
		model.Character
		Skills []model.Skill `json:"skills"`
	}  `json:"characters"`
}

func HandleUser(conn websocket.Connection, user model.User) {
	log := func(args ...interface{}) {
		fmt.Printf("(%s) [%s] %s: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"), user.Name)
		fmt.Println(args...)
	}
	log("login")
	data := DataPack{}
	db.Model(&model.Character{}).Where("user_id = ?", user.ID).Scan(&data.Characters)

	//conn.On("GAME_GET_BAG", func() {}) todo
	conn.On("GAME_GET_PARTY", func(town_id int) {
		conn.Emit("GAME_GET_PARTY_SUCCESS", data)
	})
	conn.OnDisconnect(func() {
		log(user.ID, " disconnect ...")
	})
}
