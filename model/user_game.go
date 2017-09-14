package model

import (
	"wandering-server/model/math"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"-";gorm:"not null;unique"`    // 用户名
	Password string `json:"-";gorm:"not null"`           // 密码
	Name     string `json:"name";gorm:"not null;unique"` // 昵称

	//BagID  int `json:"bag_id"`
	//Active int `json:"active"`
	//Mental int `json:"mental"`
	//ActiveLimit int `json:"active_limit";gorm:"-"`
	//MentalLimit int `json:"mental_limit";gorm:"-"`
}

type Game struct {
	ID     int `json:"id"`      // 游戏ID
	UserID int `json:"user_id"` // 用户ID

	Name     string     `json:"name"`     // 名称
	Position math.Point `json:"position"` // 地图位置
	MapID    int        `json:"-"`   // 地图ID
	BagID    int        `json:"-"`   // 背包ID
	PartyID  int        `json:"-"` // 队伍ID
}

func GetUserByAuthenticate(username string, password string) (User, error) {
	user := User{} // 获取user
	err := DB.Model(&User{}).Where(User{Username: username, Password: password}).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}
