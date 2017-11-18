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
	DB.AutoMigrate(
		&Skill{},
		&Bag{},
		&User{},
		&Game{},
		&Item{},
	)

	fmt.Println("Model init success..")
}
