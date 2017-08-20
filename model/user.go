package model

type User struct {
	Username string
	Password string
}

type UserAccess struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
