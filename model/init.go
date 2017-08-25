package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:2101429@tcp(127.0.0.1:3306)/wandering?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(
		&Character{},
		&Skill{},
		&User{},
		&Town{},
		&Tachie{},
	)
	// Skills 外键 Skill.CharID -> Character.ID
	//DB.Model(&Skill{}).AddForeignKey("char_id", "characters(id)", "CASCADE", "CASCADE")
	//DB.Model(&Character{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	fmt.Println("Model init success..")
}
