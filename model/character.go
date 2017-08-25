package model

type Character struct {
	ID     int
	UserID int    `gorm:"index"`
	Name   string `gorm:"not null"`
	TownID string `gorm:"index"`
}
