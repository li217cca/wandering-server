package model

import "time"

type History struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"time"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
}
