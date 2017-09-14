package handler

import (
	"wandering-server/common"
	"image"
)

func (this *Context) OnHandleCard(ID int)  {
	card, err := this.Scene.GetCard(ID)
	if err != nil {
		this.Emit(common.GAME_RECEIPT_MESSAGE, err)
		return
	}
	card.OnHandle()
}

func (this *Context) OnHandleAheadCard() {
	this.User.Position.Add(image.Point{10, 10,})
	this.EmitUser()
	this.GenerateScene()
}
func (this *Context) OnHandleBackCard() {
	this.User.Position.Add(image.Point{-10, -10,})
	this.EmitUser()
	this.GenerateScene()
}