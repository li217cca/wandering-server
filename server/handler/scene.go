package handler

import (
	"wandering-server/common"
)

func (this *Context) GenerateScene()  {

	if this.Time.Weather == common.WEATHER_AUTUMN {
		this.Scene.AddCards(Card{
			Name: "秋天",
			Type: "weather",
			OnHandle: func() {
				this.User.Mental -= 5
				this.EmitUser()
			},
		})
	}

	this.Scene.AddCards(Card{
		Name: "后退",
		Type: "action",
		OnHandle: this.OnHandleBackCard,
	}, Card{
		Name: "前进",
		Type: "action",
		OnHandle: this.OnHandleAheadCard,
	})
	this.EmitScene()
}