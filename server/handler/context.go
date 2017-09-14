package handler

import (
	"fmt"
	"github.com/kataras/iris/websocket"
	"time"
	"wandering-server/common"
	"wandering-server/model"
)

type Context struct {
	Conn websocket.Connection `json:"-"`
	User model.User           `json:"user"`
	Game Game `json:"game"`
	Log  func(...interface{})
	Emit func(string, interface{}) error
}
type Game struct {
	//Party model.Party `json:"party"`
	//Bag   model.Bag         `json:"bag"`
	//Scene Scene             `json:"scene"`
	Time  *model.Time       `json:"time"`
}

// 新建游戏内容实体
func NewUserContext(user model.User, conn websocket.Connection, nowTime *model.Time) Context {
	logger := func(args ...interface{}) {
		fmt.Printf("(%s) [%s] %s: ", conn.Context().RemoteAddr(), time.Now().Format("01/02 15:04:05.00"), user.Name)
		fmt.Println(args...)
	}
	ctx := Context{
		conn, // socket 连接
		user, // 用户信息
		Game{ // 游戏信息
			//model.GetPartyByID(user.Game.PartyID), // 队伍信息
			//model.GetBagByID(user.Game.BagID),   // 背包信息
			//Scene: Scene{},                         // 游戏场景
			nowTime,                         // 游戏内时间
		},
		logger,    // 日志器
		conn.Emit, // 客户端通信
	}
	ctx.RefreshState()
	return ctx
}

// 刷新状态
func (this *Context) RefreshState() {
	// todo..
}

func (this *Context) EmitParty() {
	this.Emit(common.GAME_RECEIPT_PARTY, this.Game.Party)
}
func (this *Context) EmitScene() {
	this.Emit(common.GAME_RECEIPT_SCENE, this.Game.Scene.GetScene())
}
func (this *Context) EmitTime() {
	this.Emit(common.GAME_RECEIPT_TIME, this.Game.Time)
}
func (this *Context) EmitBag() {
	this.Emit(common.GAME_RECEIPT_BAG, this.Game.Bag)
}
func (this *Context) EmitUser() {
	this.Emit(common.GAME_RECEIPT_USER, this.User)
}
func (this *Context) EmitMessage(i ...interface{}) {
	this.Emit(common.GAME_RECEIPT_MESSAGE, i)
}
