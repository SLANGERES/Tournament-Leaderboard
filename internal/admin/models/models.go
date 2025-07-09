package models

type CreateAdmin struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LogInAdmin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
