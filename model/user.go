package model

import "image"

type User struct {
	ID       int `json:"id"`
	Name     string `json:"name";gorm:"not null"`
	Username string `json:"username";gorm:"not null"`
	Password string `json:"password";gorm:"not null"`
	Position image.Point `json:"position"`
}
