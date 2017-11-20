package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgreSQL needs
)

/*
DB: the database
*/
var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres",
		"host=localhost port=5432 user=wandering dbname=wandering sslmode=disable password=2101429")
	if err != nil {
		panic(err)
	}
	if err = DB.AutoMigrate(
		&Skill{},
		&Bag{},
		&Item{},
		&User{},
		&Game{},
		&Charactor{},
		&Map{},
		&Resource{},
		&Route{},
	).Error; err != nil {
		fmt.Printf("model.init() error 01 %v\n", err)
	}

	fmt.Println("Model init success..")
}
