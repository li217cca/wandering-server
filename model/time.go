package model

type Time struct {
	Year int `json:"year"`
	Weather string `json:"weather"`
	Day int `json:"day"`
}

const (
	WEATHER_SPRING = "spring"
	WEATHER_SUMMER = "summer"
	WEATHER_AUTUMN = "autumn"
	WEATHER_WINTER = "winter"
)