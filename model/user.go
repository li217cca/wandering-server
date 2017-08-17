package model

type User struct {
	Name     string
	Username string
	Password string
}

type UserAccess struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
