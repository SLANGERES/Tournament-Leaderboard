package models

type CreateUser struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
