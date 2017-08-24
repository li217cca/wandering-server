package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db, err := gorm.Open("mysql", "root:2101429@tcp(127.0.0.1:3306)/wandering?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("model init..", db.HasTable("characters"))
}