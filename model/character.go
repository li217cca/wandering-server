package model

type Character struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id";gorm:"index"`
	Name   string `json:"name";gorm:"not null"`
	TownID string `json:"town_id";gorm:"index"`
}
