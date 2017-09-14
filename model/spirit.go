package model

type Spirit struct {
	ID  int    `json:"id"`
	URL string `json:"url"`

	Width  int `json:"width"`
	Height int `json:"height"`
}
