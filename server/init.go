package server

import (
	"github.com/jinzhu/gorm"
	"wandering-server/model"
)

var (
	db *gorm.DB
)

func init() {
	db = model.DB
}
