package model

import (
	"fmt"
)

/*
User user struct
*/
type User struct {
	ID       int    `json:"id"`
	Username string `json:"-" gorm:"not null;unique"` // 用户名
	Password string `json:"-" gorm:"not null"`        // 密码
}

/*
GetUserByAuthenticate ...
*/
func GetUserByAuthenticate(username string, password string) (User, error) {
	user := User{} // 获取user
	err := DB.Model(&User{}).Where(User{Username: username, Password: password}).Find(&user).Error
	if err != nil {
		return user, fmt.Errorf("\nGetUserByAuthenticate 01 \n%v", err)
	}
	return user, nil
}
