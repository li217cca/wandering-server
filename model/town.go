package model

type Town struct {
	ID     int
	Name   string `gorm:"not null"`
}