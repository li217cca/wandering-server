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
	conn.On("GAME_GET_TIME", func(opt int) {
		conn.Emit("GAME_GET_TIME_SUCCESS", nowTime)
	})
	scene := model.Scene{
		Cards: []model.Card{
			model.Card{
				"后退",
				0,
				"action",
				0,
			},
			model.Card{
				"前进",
				0,
				"action",
				0,
			},
		},
	}
	conn.On("GAME_GET_SCENE", func(opt int) {
		log("GAME_GET_SCENE", scene)
		conn.Emit("GAME_GET_SCENE_SUCCESS", scene)
	})
	conn.On("GAME_HANDLE_CARD", func(card interface{}) {
		log("HANDLE CARD")
		cardMap := card.(map[string]interface{})
		scene.Cards = append(scene.Cards, model.Card{
			Name: cardMap["name"].(string),
			Type: cardMap["type"].(string),
		})
		conn.Emit("GAME_GET_SCENE_SUCCESS", scene)
	})
	conn.OnDisconnect(func() {
		log(user.ID, " disconnect ...")
	})
}
